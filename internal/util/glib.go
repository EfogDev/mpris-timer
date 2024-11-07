package util

import (
	"context"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/efogdev/gotk4-adwaita/pkg/adw"
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
