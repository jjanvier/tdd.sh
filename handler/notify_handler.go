package handler

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/gen2brain/beeep"
	"os"
	"path/filepath"
	"time"
)

const title = "TDD.sh"
const icon = "assets/logo.png"
const bell = "assets/bell.ogg"

type NotifyHandlerI interface {
	HandleNotify(delay int, message string) error
}

type NotifyHandler struct{}

func (handler NotifyHandler) HandleNotify(delay int, message string) error {
	wait(delay)

	iconPath, _ := filepath.Abs(icon)
	displayNotification(message, iconPath)

	bellPath, _ := filepath.Abs(bell)
	playBell(bellPath)

	return nil
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
