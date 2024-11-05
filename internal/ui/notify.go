package ui

import (
	"context"
	_ "embed"
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/google/uuid"
	"log"
)

//go:embed icon.svg
var icon string

func Notify(title string, text string) {
	log.Printf("notify: %s", title)

	// a workaround; could've used notify-send
	// but then default click action is to open timer again
	// which is not desired
	nApp := adw.NewApplication("io.github.efogdev.mpris-timer", gio.ApplicationNonUnique)
	err := nApp.Register(context.Background())
	if err != nil {
		log.Printf("error registering application: %v", err)
	}

	nApp.ConnectActivate(func() {
		id, _ := uuid.NewV7()
		actionName := "app." + id.String()
		nApp.AddAction(gio.NewSimpleAction(actionName, nil))

		n := gio.NewNotification(title)
		n.SetBody(text)
		n.SetPriority(gio.NotificationPriorityUrgent)
		n.SetDefaultAction(actionName)
		n.SetIcon(gio.NewBytesIcon(glib.NewBytes([]byte(icon))))

		nApp.SendNotification(id.String(), n)
	})

	nApp.Run(nil)
}
