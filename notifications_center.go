package main

import (
	"github.com/gen2brain/beeep"
	"path/filepath"
)

const title = "TDD.sh"
const icon = "assets/logo.png"

type NotificationCenterI interface {
	Notify(message string)
	Alert(message string)
}

type NotificationsCenter struct {
}

func (center NotificationsCenter) Notify(message string) {
	iconPath, _ := filepath.Abs(icon)
	beeep.Notify(title, message, iconPath)
}

func (center NotificationsCenter) Alert(message string) {
	iconPath, _ := filepath.Abs(icon)
	beeep.Alert(title, message, iconPath)
}
