package main

import (
	"os"
	"strconv"
)

const notificationPackage = "notification"

type NotificationCenterI interface {
	NotifyWithDelay(alias string, delay int, message string)
	Reset(alias string)
}

type NotificationsCenter struct {
	executor CommandExecutorI
	pidFileName string
}

func (center NotificationsCenter) NotifyWithDelay(alias string, delay int, message string) {
	cmd := Command{notificationPackage, []string{strconv.Itoa(delay), message}}
	pid, _ := center.executor.ExecuteBackground(cmd)

	timers := LoadTimers(center.pidFileName)
	timers.UpsertPid(alias, pid)

	SaveTimers(center.pidFileName, timers)
}

func (center NotificationsCenter) Reset(alias string) {
	timers := LoadTimers(center.pidFileName)
	pid := timers.GetPid(alias)

	killPreviousNotification(pid)
}

func killPreviousNotification(pid int) {
	process, err := os.FindProcess(pid)
	if err == nil {
		process.Kill()
	}
}
