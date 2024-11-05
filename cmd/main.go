package main

import (
	"context"
	"flag"
	"log"
	"mpris-timer/internal/player"
	"mpris-timer/internal/ui"
	"os"
	"os/signal"
)

var (
	notify   bool
	sound    bool
	useUI    bool
	duration int
	title    string
	text     string
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	flag.BoolVar(&notify, "notify", true, "Send desktop notification")
	flag.BoolVar(&sound, "sound", true, "Play sound")
	flag.BoolVar(&useUI, "ui", false, "Show timepicker UI (default true)")
	flag.IntVar(&duration, "start", 0, "Start the timer immediately")
	flag.StringVar(&title, "title", "Timer", "Name/title of the timer")
	flag.StringVar(&text, "text", "Time is up!", "Notification text")
	flag.Parse()

	if useUI && duration > 0 {
		log.Fatalf("UI can't be used with -start")
	}

	if !useUI && duration == 0 {
		useUI = true
	}

	if useUI {
		log.Println("UI launched")
		ui.Init(&duration)
	}

	log.Printf("timer started: %d sec", duration)
	timer, err := player.NewMPRISPlayer(duration, title)
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
		if notify {
			log.Printf("desktop notification requested")
			ui.Notify(timer.Name, text)
		}

		// note: synchronous
		if sound {
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
