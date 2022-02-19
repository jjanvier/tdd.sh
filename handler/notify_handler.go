package handler

import (
	"embed"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/gen2brain/beeep"
	"io/fs"
	"io/ioutil"
	"os"
	"time"
)

const title = "TDD.sh"
const icon = "assets/logo.png"
const bell = "assets/bell.ogg"

// assets present in assets/* are embedded in a local filesystem
// see https://pkg.go.dev/embed@master and main.go
var Assets embed.FS

type NotifyHandlerI interface {
	HandleNotify(delay int, message string) error
}

type NotifyHandler struct{}

func (handler NotifyHandler) HandleNotify(delay int, message string) error {
	wait(delay)

	iconData, _ := Assets.ReadFile(icon)
	displayNotification(message, iconData)

	bellFile, _ := Assets.Open(bell)
	playBell(bellFile)

	return nil
}

func wait(delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
}

// beeep.Notify wants a path of file for the icon
// our icon file is embedded in the binary's filesystem, which means we can't access it from TDD.sh's host
// that's why we create a temporary icon file on the host
func displayNotification(message string, iconFileContent []byte) {
	tmpIconFile := createTmpIconFile(iconFileContent)
	defer tmpIconFile.Close()
	defer os.Remove(tmpIconFile.Name())

	beeep.Notify(title, message, tmpIconFile.Name())
}

// see the tutorial https://github.com/faiface/beep/wiki/Hello,-Beep! to understand how it works
func playBell(bell fs.File) {
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

func createTmpIconFile(iconContent []byte) *os.File {
	tmpFile, _ := ioutil.TempFile(os.TempDir(), "tddshicon-")
	tmpFile.Write(iconContent)

	return tmpFile
}
