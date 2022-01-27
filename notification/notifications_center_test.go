package notification

import (
	"github.com/jjanvier/tdd/execution"
	"github.com/jjanvier/tdd/helper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strconv"
	"testing"
)

func TestNotifyWithDelay(t *testing.T) {
	pidFile := helper.CreateTmpFile("")
	defer helper.RemoveTmpFile(pidFile)

	executor := execution.SuccessCommandExecutorMock{}
	center := NotificationsCenter{executor, pidFile.Name()}

	executor.On("ExecuteBackground").Once()
	center.NotifyWithDelay("ut", 45, "the message")

	actualContent, _ := ioutil.ReadFile(pidFile.Name())
	// TODO: not good, we expose some internal details here
	expectedContent := `pids:
  ut: ` + strconv.Itoa(execution.FakeCommandPid) + "\n"

	assert.Equal(t, expectedContent, string(actualContent))
}

func TestNotifyWithDelayWithAPreviousNotification(t *testing.T) {
	// a PID is already present in the file, so a notification is already scheduled
	// at the end, this PID should be replaced
	content := `pids:
  ut: 654
`

	pidFile := helper.CreateTmpFile(content)
	defer helper.RemoveTmpFile(pidFile)

	executor := execution.SuccessCommandExecutorMock{}
	center := NotificationsCenter{executor, pidFile.Name()}

	executor.On("ExecuteBackground").Once()
	center.NotifyWithDelay("ut", 45, "the message")

	actualContent, _ := ioutil.ReadFile(pidFile.Name())
	// TODO: not good, we expose some internal details here
	expectedContent := `pids:
  ut: ` + strconv.Itoa(execution.FakeCommandPid) + "\n"

	assert.Equal(t, expectedContent, string(actualContent))
}
