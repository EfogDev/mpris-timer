package util

import (
	"flag"
)

var (
	Notify   bool
	Sound    bool
	Volume   float64
	UseUI    bool
	Duration int
	Title    string
	Text     string
	Color    string
)

func LoadFlags() {
	flag.BoolVar(&Notify, "notify", UserPrefs.EnableNotification, "Send desktop notification")
	flag.BoolVar(&Sound, "sound", UserPrefs.EnableSound, "Play sound")
	flag.Float64Var(&Volume, "volume", UserPrefs.Volume, "Volume (0-1)")
	flag.BoolVar(&UseUI, "ui", false, "Show timepicker UI (default true)")
	flag.IntVar(&Duration, "start", 0, "Start the timer immediately")
	flag.StringVar(&Title, "title", UserPrefs.DefaultTitle, "Name/title of the timer")
	flag.StringVar(&Text, "text", UserPrefs.DefaultText, "Notification text")
	flag.StringVar(&Color, "color", UserPrefs.ProgressColor, "Progress color in the player, rgb hex value, i.e. #000000")
	flag.Parse()
}
