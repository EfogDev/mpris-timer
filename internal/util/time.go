package util

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func TimeFromPreset(preset string) time.Time {
	result, err := time.Parse("04:05", preset)
	if err != nil {
		log.Fatalf("failed to parse preset %s: %v", preset, err)
	}

	return result
}

func TimeFromParts(hours int, minutes int, seconds int) time.Time {
	result, err := time.Parse("15:04:05", fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds))
	if err != nil {
		log.Fatalf("failed to parse parts %d %d %d: %v", hours, minutes, seconds, err)
	}

	return result
}

func TimeFromStrings(hours string, minutes string, seconds string) time.Time {
	hoursInt, err := strconv.Atoi(hours)
	if err != nil {
		log.Fatalf("failed to parse hours %s: %v", hours, err)
	}

	minutesInt, err := strconv.Atoi(minutes)
	if err != nil {
		log.Fatalf("failed to parse minutes %s: %v", minutes, err)
	}

	secondsInt, err := strconv.Atoi(seconds)
	if err != nil {
		log.Fatalf("failed to parse seconds %s: %v", seconds, err)
	}

	return TimeFromParts(hoursInt, minutesInt, secondsInt)
}
