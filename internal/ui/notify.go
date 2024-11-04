package ui

import (
	"bytes"
	"context"
	"log"
	"os/exec"
)

func Notify(ctx context.Context, title string, text string) {
	log.Printf("notify: %s", title)

	var buf bytes.Buffer
	args := []string{"-a", title, "-i", "io.github.efogdev.mpris-timer", "-u", "critical", "-w", "-e", text, "-A", "Ok"}
	cmd := exec.CommandContext(ctx, "notify-send", args...)
	cmd.Stdout = &buf

	if err := cmd.Run(); err != nil {
		log.Fatalf("notify-send: %v", err)
	}

	log.Printf("%s", buf.String())
}
