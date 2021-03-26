package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const notificationPackage = "notification"

type NotificationCenterI interface {
	NotifyWithDelay(delay int, message string)
	Reset()
}

type NotificationsCenter struct {
	executor CommandExecutorI
	pidFileName string
}

func (center NotificationsCenter) NotifyWithDelay(delay int, message string) {
	cmd := Command{notificationPackage, []string{strconv.Itoa(delay), message}}
	pid, _ := center.executor.ExecuteBackground(cmd)
	putPidToPidFile(center.pidFileName, pid)
}

func (center NotificationsCenter) Reset() {
	killPreviousNotification(center.pidFileName)
}

func killPreviousNotification(filename string) {
	pid, err := getPidFromPidFile(filename)
	if err != nil {
		return
	}

	process, err := os.FindProcess(pid)
	if err == nil {
		process.Kill()
	}
}

func putPidToPidFile(filename string, pid int) {
	pidFile, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	defer pidFile.Close()

	pidFile.WriteString(strconv.Itoa(pid) + "\n")
}

func getPidFromPidFile(filename string) (int, error) {
	pidByte, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	pid, err := strconv.Atoi(strings.TrimSpace(string(pidByte)))
	if err != nil {
		return 0, err
	}

	return pid, nil
}
