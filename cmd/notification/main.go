package main

import (
	"github.com/gen2brain/beeep"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const title = "TDD.sh"
const icon = "assets/logo.png"

func main() {
	delay, _ := strconv.Atoi(os.Args[1])
	message := os.Args[2]

	time.Sleep(time.Duration(delay) * time.Second)

	iconPath, _ := filepath.Abs(icon)
	beeep.Notify(title, message, iconPath)
}
