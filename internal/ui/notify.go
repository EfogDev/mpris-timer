package ui

import (
	"context"
	_ "embed"
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/google/uuid"
	"log"
	"mpris-timer/internal/util"
)

//go:embed res/icon.svg
var icon []byte

func Notify(title string, text string) {
	log.Printf("notify: %s", title)

	// a workaround; could've used notify-send
	// but then default click action is to open timer again
	// which is not desired
	if !util.UseUI {
		sendNotification(util.App, title, text)
	} else {
		nApp := adw.NewApplication(util.AppId, gio.ApplicationNonUnique)
		nApp.ConnectActivate(func() {
			sendNotification(nApp, title, text)
		})

		_ = nApp.Register(context.Background())
		nApp.Run(nil)
	}
}

func sendNotification(app *adw.Application, title string, text string) {
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
