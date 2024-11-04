package util

const (
	width         = 256
	height        = 256
	padding       = 16
	strokeWidth   = 32
	fgStrokeColor = "#535353"
	bgStrokeColor = "#2190a4"

	DefaultPreset = "02:30"
)

var DefaultPresets = []string{
	"00:30",
	"01:00",
	"01:30",
	"02:00",
	"02:30",
	"03:00",
	"05:00",
	"07:00",
	"10:00",
	"15:00",
	"20:00",
	"30:00",
}

const svgTemplate = `
<svg width="{{.Width}}" height="{{.Height}}">
    <circle cx="{{.CenterX}}" cy="{{.CenterY}}" r="{{.Radius}}" fill="none" stroke="{{.FgStrokeColor}}" stroke-width="{{.BaseWidth}}" />
    <circle cx="{{.CenterX}}" cy="{{.CenterY}}" r="{{.Radius}}" fill="none" stroke="{{.BgStrokeColor}}" stroke-width="{{.StrokeWidth}}" stroke-dasharray="{{.Circumference}}" stroke-dashoffset="{{.DashOffset}}" transform="rotate(-90 {{.CenterX}} {{.CenterY}})" />
</svg>
`

type svgParams struct {
	Width         int
	Height        int
	CenterX       int
	CenterY       int
	Radius        float64
	FgStrokeColor string
	BaseWidth     int
	BgStrokeColor string
	StrokeWidth   int
	Circumference float64
	DashOffset    float64
}
