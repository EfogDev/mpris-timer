package util

import (
	"flag"
)

var (
	Notify   bool
	Sound    bool
	Volume   float64
	Silence  int
	UseUI    bool
	Duration int
	Title    string
	Text     string
	Color    string
)

func LoadFlags() {
	flag.BoolVar(&Notify, "notify", UserPrefs.EnableNotification, "Send desktop notification")
	flag.BoolVar(&Sound, "sound", UserPrefs.EnableSound, "Play sound")
	flag.Float64Var(&Volume, "volume", UserPrefs.Volume, "Volume [0-1]")
	flag.IntVar(&Silence, "silence", 0, "Play this milliseconds of silence before the actual audio — might be helpful for audio devices that wake up not immediately")
	flag.BoolVar(&UseUI, "ui", false, "Show timepicker UI (default true)")
	flag.IntVar(&Duration, "start", 0, "Start the timer immediately")
	flag.StringVar(&Title, "title", UserPrefs.DefaultTitle, "Name/title of the timer")
	flag.StringVar(&Text, "text", UserPrefs.DefaultText, "Notification text")
	flag.StringVar(&Color, "color", UserPrefs.ProgressColor, "Progress color for the player")
	flag.Parse()
}
