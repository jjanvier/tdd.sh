package handler

import (
	"embed"
	"errors"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/gen2brain/beeep"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

const title = "TDD.sh"
const icon = "assets/logo.png"
const localIcon = "tddsh-icon.png"
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
	localIconFile := getLocalIconFile(iconData)
	displayNotification(message, localIconFile)

	bellFile, _ := Assets.Open(bell)
	playBell(bellFile)

	return nil
}

func wait(delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
}

func displayNotification(message string, icon *os.File) {
	beeep.Notify(title, message, icon.Name())
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

// beeep.Notify wants a path of file for the icon
// our icon file is embedded in the binary's filesystem, which means we can't access it from TDD.sh's host
// that's why we create an icon file on the host
func getLocalIconFile(iconContent []byte) *os.File {
	if shouldCreateLocalIconFile() {
		createLocalIconFile(iconContent)
	}

	file, _ := os.OpenFile(getLocalIconFilePath(), os.O_RDONLY, 0644)

	return file
}

func getLocalIconFilePath() string {
	path := os.TempDir() + "/" + localIcon
	realPath, _ := filepath.Abs(filepath.FromSlash(path))

	return realPath
}

func shouldCreateLocalIconFile() bool {
	if _, err := os.Stat(getLocalIconFilePath()); errors.Is(err, os.ErrNotExist) {
		return true
	}

	return false
}

func createLocalIconFile(iconContent []byte) {
	file, _ := os.OpenFile(getLocalIconFilePath(), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	file.Write(iconContent)
	defer file.Close()
}
