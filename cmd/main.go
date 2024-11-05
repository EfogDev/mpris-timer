package main

import (
	"context"
	"log"
	"mpris-timer/internal/player"
	"mpris-timer/internal/ui"
	"mpris-timer/internal/util"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	util.RegisterApp(ctx)
	util.LoadPrefs()
	util.LoadFlags()

	if util.UseUI && util.Duration > 0 {
		log.Fatalf("UI can't be used with -start")
	}

	if !util.UseUI && util.Duration == 0 {
		util.UseUI = true
	}

	if util.UseUI {
		log.Println("UI launched")
		ui.Init()
	}

	log.Printf("timer started: %d sec", util.Duration)
	timer, err := player.NewMPRISPlayer(util.Duration, util.Title)
	if err != nil {
		log.Fatalf("failed to create player: %v", err)
	}

	if err := timer.Start(); err != nil {
		log.Fatalf("failed to start timer: %v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	select {
	case <-timer.Done:
		log.Println("timer done")

		// note: synchronous
		if util.Notify {
			log.Printf("desktop notification requested")
			ui.Notify(timer.Name, util.Text)
		}

		// note: synchronous
		if util.Sound {
			log.Printf("sound requested")
			err = ui.PlayAudio()
			if err != nil {
				log.Printf("error playing sound file: %v", err)
			}
		}

		cancel()
	case <-sigChan:
		log.Println("interrupt received")
		cancel()
	case <-ctx.Done():
		timer.Destroy()
		log.Println("context done")
		return
	}
}
