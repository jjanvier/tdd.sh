package notification

import (
	"github.com/jjanvier/tdd/helper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestLoadTimersFromPidFile(t *testing.T) {
	content := `pids:
  ut: 655
  another_alias: 987976
`
	pidFile := helper.CreateTmpFile(content)
	defer helper.RemoveTmpFile(pidFile)

	actual := LoadTimers(pidFile.Name())

	expected := Timers{}
	aliases := make(map[string]int)
	aliases["ut"] = 655
	aliases["another_alias"] = 987976
	expected.Pids = aliases

	assert.Equal(t, expected, actual)
}

func TestSaveTimersInPidFile(t *testing.T) {
	pidFile := helper.CreateTmpFile("")
	defer helper.RemoveTmpFile(pidFile)

	timers := Timers{}
	aliases := make(map[string]int)
	aliases["bar"] = 9875421
	aliases["foo"] = 4123
	timers.Pids = aliases

	SaveTimers(pidFile.Name(), timers)

	actual, _ := ioutil.ReadFile(pidFile.Name())
	expected := `pids:
  bar: 9875421
  foo: 4123
`
	assert.Equal(t, expected, string(actual))
}

func TestSaveTimersInNonEmptyPidFile(t *testing.T) {
	content := `pids:
  just_an_alias: 655
  baz: 987976
`

	pidFile := helper.CreateTmpFile(content)
	defer helper.RemoveTmpFile(pidFile)

	timers := Timers{}
	aliases := make(map[string]int)
	aliases["bar"] = 12
	aliases["foo"] = 4123
	timers.Pids = aliases

	SaveTimers(pidFile.Name(), timers)

	actual, _ := ioutil.ReadFile(pidFile.Name())
	expected := `pids:
  bar: 12
  foo: 4123
`
	assert.Equal(t, expected, string(actual))
}

func TestItGetsThePidOfATimer(t *testing.T) {
	timers := Timers{}
	aliases := make(map[string]int)
	aliases["bar"] = 12
	aliases["foo"] = 654
	timers.Pids = aliases

	actual := timers.GetPid("foo")
	assert.Equal(t, 654, actual)
}

func TestItInsertsThePidOfATimer(t *testing.T) {
	timers := Timers{}
	timers.UpsertPid("foo", 545)

	assert.Equal(t, 545, timers.Pids["foo"])
}

func TestItReplacesThePidOfATimer(t *testing.T) {
	timers := Timers{}
	aliases := make(map[string]int)
	aliases["bar"] = 9797
	aliases["foo"] = 654
	timers.Pids = aliases
	timers.UpsertPid("bar", 18)

	assert.Equal(t, 18, timers.Pids["bar"])
}
