package notification

import (
	"github.com/jjanvier/tdd/execution"
	"os"
	"strconv"
)

const notificationPackage = "notification"

type NotificationCenterI interface {
	NotifyWithDelay(alias string, delay int, message string)
	Reset(alias string)
}

type NotificationsCenter struct {
	Executor    execution.CommandExecutorI
	PidFileName string
}

func (center NotificationsCenter) NotifyWithDelay(alias string, delay int, message string) {
	cmd := execution.Command{Name: notificationPackage, Arguments: []string{strconv.Itoa(delay), message}}
	pid, _ := center.Executor.ExecuteBackground(cmd)

	timers := LoadTimers(center.PidFileName)
	timers.UpsertPid(alias, pid)

	SaveTimers(center.PidFileName, timers)
}

func (center NotificationsCenter) Reset(alias string) {
	timers := LoadTimers(center.PidFileName)
	pid := timers.GetPid(alias)

	killPreviousNotification(pid)
}

func killPreviousNotification(pid int) {
	process, err := os.FindProcess(pid)
	if err == nil {
		process.Kill()
	}
}
