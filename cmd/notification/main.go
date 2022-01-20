package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/gen2brain/beeep"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const title = "TDD.sh"
const icon = "assets/logo.png"
const bell = "assets/bell.ogg"

/**
 This script display a notification message and plays a bell after a delay.

 Usage:
	go run cmd/notification/main.go delay message

The notification is displayed thanks to https://github.com/gen2brain/beeep.
The bell is played thanks to https://github.com/faiface/beep. Beep requires Oto, which itself requires libasound2-dev on Linux.
*/
func main() {
	delay, _ := strconv.Atoi(os.Args[1])
	message := os.Args[2]

	wait(delay)

	iconPath, _ := filepath.Abs(icon)
	displayNotification(message, iconPath)

	bellPath, _ := filepath.Abs(bell)
	playBell(bellPath)
}

func wait(delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
}

func displayNotification(message string, iconPath string) {
	beeep.Notify(title, message, iconPath)
}

// see the tutorial https://github.com/faiface/beep/wiki/Hello,-Beep! to understand how it works
func playBell(bellPath string) {
	bell, _ := os.Open(bellPath)
	defer bell.Close()

	streamer, format, _ := vorbis.Decode(bell)
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
