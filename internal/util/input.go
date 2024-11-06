package util

import (
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"strings"
)

func ParseKeyval(keyval uint) string {
	return strings.ReplaceAll(gdk.KeyvalName(keyval), "KP_", "")
}

func IsGdkKeyvalNumber(keyval uint) bool {
	return (keyval >= gdk.KEY_0 && keyval <= gdk.KEY_9) || (keyval >= gdk.KEY_KP_0 && keyval <= gdk.KEY_KP_9)
}
