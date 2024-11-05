package util

import (
	"flag"
)

var (
	Notify   bool
	Sound    bool
	UseUI    bool
	Duration int
	Title    string
	Text     string
)

func LoadFlags() {
	flag.BoolVar(&Notify, "notify", UserPrefs.EnableNotification, "Send desktop notification")
	flag.BoolVar(&Sound, "sound", UserPrefs.EnableSound, "Play sound")
	flag.BoolVar(&UseUI, "ui", false, "Show timepicker UI (default true)")
	flag.IntVar(&Duration, "start", 0, "Start the timer immediately")
	flag.StringVar(&Title, "title", UserPrefs.DefaultTitle, "Name/title of the timer")
	flag.StringVar(&Text, "text", UserPrefs.DefaultText, "Notification text")
	flag.Parse()
}
