package util

import (
	"fmt"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"log"
	"strings"
	"time"
)

func ParseKeyval(keyval uint) string {
	return strings.ReplaceAll(gdk.KeyvalName(keyval), "KP_", "")
}

func IsGdkKeyvalNumber(keyval uint) bool {
	return (keyval >= gdk.KEY_0 && keyval <= gdk.KEY_9) || (keyval >= gdk.KEY_KP_0 && keyval <= gdk.KEY_KP_9)
}

func NumToLabelText(num int) string {
	if num > 59 || num < 0 {
		log.Fatalf("NumToLabelText: num must be between 0 and 59")
	}

	return fmt.Sprintf("%02d", num)
}

func FormatDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second

	if h > 0 {
		return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
	}

	return fmt.Sprintf("%02d:%02d", m, s)
}
