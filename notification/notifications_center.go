package notification

import (
	"github.com/jjanvier/tdd/execution"
	"os"
)

type NotificationCenterI interface {
	// NotifyWithDelay launch a notification after a delay, and in background, for the given alias
	NotifyWithDelay(alias string, delay int, message string)
	// Reset cancel the previous notification that was launched in background, if any, for the given alias
	Reset(alias string)
}

// NotificationsCenter allows handling notifications. Internally, the PidFileName is used to store the PIDs of the notifications
// related to the different test aliases that have been launched. See Timers
type NotificationsCenter struct {
	Executor       execution.CommandExecutorI
	CommandFactory execution.CommandFactory
	PidFileName    string
}

func (center NotificationsCenter) NotifyWithDelay(alias string, delay int, message string) {
	cmd := center.CommandFactory.CreateNotify(delay, message)

	// call this command in a subprocess as we don't want the current
	// process (our "tdd launch myalias" command) to wait for the delay
	pid, err := center.Executor.ExecuteBackground(cmd)

	if err == nil {
		saveCommandPid(center.PidFileName, alias, pid)
	}
}

func (center NotificationsCenter) Reset(alias string) {
	timers := LoadTimers(center.PidFileName)
	pid, err := timers.GetPid(alias)

	if err == nil {
		killPreviousCommand(pid)
	}
}

func saveCommandPid(pidFileName string, alias string, pid int) {
	timers := LoadTimers(pidFileName)
	timers.UpsertPid(alias, pid)
	SaveTimers(pidFileName, timers)
}

func killPreviousCommand(pid int) {
	process, err := os.FindProcess(pid)
	if err == nil {
		process.Kill()
	}
}
