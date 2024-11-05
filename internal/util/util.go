package util

import (
	"bytes"
	"fmt"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"math"
	"os"
	"path"
	"strings"
	"text/template"
)

var (
	CacheDir string
	DataDir  string
)

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

func MakeProgressCircle(progress float64) (string, error) {
	progress = math.Max(0, math.Min(100, progress))
	filename := path.Join(CacheDir, fmt.Sprintf("_f4g.%.1f.svg", progress))

	if _, err := os.Stat(filename); err == nil {
		return filename, nil
	}

	centerX := width / 2
	centerY := height / 2
	radius := float64(width)/2 - float64(strokeWidth) - float64(padding)
	baseWidth := int(math.Round(strokeWidth * 0.25))
	circumference := 2 * math.Pi * radius
	dashOffset := circumference * (1 - progress/100)

	data := svgParams{
		Width:         width,
		Height:        height,
		CenterX:       centerX,
		CenterY:       centerY,
		Radius:        radius,
		BaseWidth:     baseWidth,
		StrokeWidth:   strokeWidth,
		FgStrokeColor: fgStrokeColor,
		BgStrokeColor: bgStrokeColor,
		Circumference: circumference,
		DashOffset:    dashOffset,
	}

	tmpl, err := template.New("svg").Parse(svgTemplate)
	if err != nil {
		return "", err
	}

	var svgBuffer bytes.Buffer
	err = tmpl.Execute(&svgBuffer, data)
	if err != nil {
		return "", err
	}

	err = os.WriteFile(filename, svgBuffer.Bytes(), 0644)
	if err != nil {
		return "", fmt.Errorf("failed to write SVG file: %w", err)
	}

	return filename, nil
}
