package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strconv"
	"testing"
)

func TestNotifyWithDelay(t *testing.T) {
	pidFile := createTmpFile("")
	defer removeTmpFile(pidFile)

	executor := successCommandExecutorMock{}
	center := NotificationsCenter{executor, pidFile.Name()}

	executor.On("ExecuteBackground").Once()
	center.NotifyWithDelay("ut", 45, "the message")

	actualContent, _ := ioutil.ReadFile(pidFile.Name())
	// TODO: not good, we expose some internal details here
	expectedContent := `pids:
  ut: ` + strconv.Itoa(commandPid) + "\n"

	assert.Equal(t, expectedContent, string(actualContent))
}

func TestNotifyWithDelayWithAPreviousNotification(t *testing.T) {
	// a PID is already present in the file, so a notification is already scheduled
	// at the end, this PID should be replaced
	content := `pids:
  ut: 654
`

	pidFile := createTmpFile(content)
	defer removeTmpFile(pidFile)

	executor := successCommandExecutorMock{}
	center := NotificationsCenter{executor, pidFile.Name()}

	executor.On("ExecuteBackground").Once()
	center.NotifyWithDelay("ut", 45, "the message")

	actualContent, _ := ioutil.ReadFile(pidFile.Name())
	// TODO: not good, we expose some internal details here
	expectedContent := `pids:
  ut: ` + strconv.Itoa(commandPid) + "\n"

	assert.Equal(t, expectedContent, string(actualContent))
}
