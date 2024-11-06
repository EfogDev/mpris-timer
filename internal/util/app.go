package util

import (
	"context"
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"log"
)

var App *adw.Application

// RegisterApp must be called before init
func RegisterApp(ctx context.Context) {
	App = adw.NewApplication(AppId, gio.ApplicationNonUnique)
	err := App.Register(ctx)
	if err != nil {
		log.Printf("error registering application: %v", err)
	}
}
