package ui

import (
	"bytes"
	"context"
	_ "embed"
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/google/uuid"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
	"log"
	"mpris-timer/internal/util"
	"time"
)

//go:embed icon.svg
var icon []byte

//go:embed ding.mp3
var sound []byte

func Notify(title string, text string) {
	log.Printf("notify: %s", title)

	// a workaround; could've used notify-send
	// but then default click action is to open timer again
	// which is not desired
	if !util.UseUI {
		SendNotification(util.App, title, text)
	} else {
		nApp := adw.NewApplication(util.AppId, gio.ApplicationNonUnique)
		nApp.ConnectActivate(func() {
			SendNotification(nApp, title, text)
		})

		_ = nApp.Register(context.Background())
		nApp.Run(nil)
	}
}

func SendNotification(app *adw.Application, title string, text string) {
	id, _ := uuid.NewV7()
	actionName := "app." + id.String()
	app.AddAction(gio.NewSimpleAction(actionName, nil))

	n := gio.NewNotification(title)
	n.SetBody(text)
	n.SetPriority(gio.NotificationPriorityUrgent)
	n.SetDefaultAction(actionName)
	n.SetIcon(gio.NewBytesIcon(glib.NewBytes(icon)))

	app.SendNotification(id.String(), n)
}

func PlayAudio() error {
	dec, err := mp3.NewDecoder(bytes.NewReader(sound))
	if err != nil {
		return err
	}

	ctx, ready, err := oto.NewContext(dec.SampleRate(), 2, 2)
	if err != nil {
		return err
	}
	<-ready

	player := ctx.NewPlayer(dec)
	defer func() { _ = player.Close() }()
	player.Play()

	for player.IsPlaying() {
		time.Sleep(time.Millisecond * 10)
	}

	return nil
}
