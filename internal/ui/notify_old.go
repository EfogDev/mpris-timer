//go:build notify_send
// +build notify_send

package ui

import (
	"bytes"
	"context"
	"log"
	"os/exec"
)

func Notify(title string, text string) {
	log.Printf("notify: %s", title)

	var buf bytes.Buffer
	args := []string{"-a", "MPRIS Timer", "-i", "io.github.efogdev.mpris-timer", "-u", "critical", "-e", title, text}
	cmd := exec.CommandContext(context.Background(), "notify-send", args...)
	cmd.Stdout = &buf

	if err := cmd.Run(); err != nil {
		log.Fatalf("notify-send: %v", err)
	}
}
