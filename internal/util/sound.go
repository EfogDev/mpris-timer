package util

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
	"log"
	"time"
)

//go:embed res/ding.mp3
var sound []byte

func PlaySound() error {
	dec, err := mp3.NewDecoder(bytes.NewReader(sound))
	if err != nil {
		return err
	}

	ctx, ready, err := oto.NewContext(dec.SampleRate(), 2, 2)
	if err != nil {
		return err
	}
	<-ready

	if Presilence != 0 {
		log.Printf("presilence requested")
		playSilence(Presilence)
	}

	player := ctx.NewPlayer(dec)
	defer func() { _ = player.Close() }()
	player.SetVolume(Volume)
	player.Play()

	for player.IsPlaying() {
		time.Sleep(time.Millisecond * 10)
	}

	return nil
}

func playSilence(ms int) {
	sampleRate := 44100
	numSamples := sampleRate * ms / 1000
	silence := make([]byte, numSamples*2)

	ctx, ready, err := oto.NewContext(sampleRate, 2, 2)
	if err != nil {
		log.Fatal(err)
	}
	<-ready

	player := ctx.NewPlayer(bytes.NewBuffer(silence))
	defer func() { _ = player.Close() }()
	player.SetVolume(1)
	player.Play()

	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}
}
