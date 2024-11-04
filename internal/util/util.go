package util

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"path"
	"text/template"
)

var tmpDir string

func init() {
	tmpDir, _ = os.UserHomeDir()
	tmpDir = path.Join(tmpDir, ".var", ".app", "io.github.efogdev.mpris-timer", "cache")
	_ = os.MkdirAll(tmpDir, 0755)
}

func MakeProgressCircle(progress float64) (string, error) {
	progress = math.Max(0, math.Min(100, progress))
	filename := path.Join(tmpDir, fmt.Sprintf("_f4g.%.1f.svg", progress))

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
