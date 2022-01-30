package notification

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// The PID file should be like:
/*
pids:
  ut: 98745
  at: 466238
*/

type Timers struct {
	Pids map[string]int
}

func (timers Timers) GetPid(alias string) (int, error) {
	if _, ok := timers.Pids[alias]; !ok {
		return -1, errors.New("No PID for alias '" + alias + "'")
	}

	return timers.Pids[alias], nil
}

func (timers *Timers) UpsertPid(alias string, pid int) {
	if timers.Pids == nil {
		timers.Pids = make(map[string]int)
	}

	timers.Pids[alias] = pid
}

func LoadTimers(pidFilePath string) Timers {
	content, err := ioutil.ReadFile(pidFilePath)
	if err != nil {
		return Timers{}
	}

	timers := Timers{}
	yaml.Unmarshal(content, &timers)

	return timers
}

func SaveTimers(pidFilePath string, timers Timers) {
	pidFile, err := os.OpenFile(pidFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	defer pidFile.Close()

	if err == nil {
		content, _ := yaml.Marshal(timers)
		pidFile.Write(content)
	}
}
