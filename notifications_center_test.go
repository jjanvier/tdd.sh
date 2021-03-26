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
	center.NotifyWithDelay(45, "the message")

	pid, _ := ioutil.ReadFile(pidFile.Name())

	expectedPid := strconv.Itoa(commandPid) + "\n"
	actualPid := string(pid)
	assert.Equal(t, expectedPid, actualPid)
}

func TestNotifyWithDelayPreviousNotification(t *testing.T) {
	// a PID is already present in the file, so a notification is already scheduled
	// at the end, it shouldn't be anymore in the file
	pidFile := createTmpFile("654\n")
	defer removeTmpFile(pidFile)

	executor := successCommandExecutorMock{}
	center := NotificationsCenter{executor, pidFile.Name()}

	executor.On("ExecuteBackground").Once()
	center.NotifyWithDelay(45, "the message")

	pid, _ := ioutil.ReadFile(pidFile.Name())

	expectedPid := strconv.Itoa(commandPid) + "\n"
	actualPid := string(pid)
	assert.Equal(t, expectedPid, actualPid)
}
