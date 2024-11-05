package util

import (
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"os"
	"path"
	"strings"
)

const (
	AppId = "io.github.efogdev.mpris-timer"

	width         = 256
	height        = 256
	padding       = 16
	strokeWidth   = 32
	bgStrokeColor = "#535353"
)

const svgTemplate = `
<svg width="{{.Width}}" height="{{.Height}}">
    <circle cx="{{.CenterX}}" cy="{{.CenterY}}" r="{{.Radius}}" fill="none" stroke="{{.BgStrokeColor}}" stroke-width="{{.BaseWidth}}" />
    <circle cx="{{.CenterX}}" cy="{{.CenterY}}" r="{{.Radius}}" fill="none" stroke="{{.FgStrokeColor}}" stroke-width="{{.StrokeWidth}}" stroke-dasharray="{{.Circumference}}" stroke-dashoffset="{{.DashOffset}}" transform="rotate(-90 {{.CenterX}} {{.CenterY}})" />
</svg>
`

var (
	CacheDir string
	DataDir  string
)

type svgParams struct {
	Width         int
	Height        int
	CenterX       int
	CenterY       int
	Radius        float64
	FgStrokeColor string
	BgStrokeColor string
	BaseWidth     int
	StrokeWidth   int
	Circumference float64
	DashOffset    float64
}

func init() {
	DataDir = glib.GetUserDataDir()
	if !strings.Contains(DataDir, AppId) {
		DataDir = path.Join(DataDir, AppId)
	}

	CacheDir, _ = os.UserHomeDir()
	CacheDir = path.Join(CacheDir, ".var", "app", AppId, "cache")

	_ = os.MkdirAll(CacheDir, 0755)
	_ = os.MkdirAll(DataDir, 0755)
}
