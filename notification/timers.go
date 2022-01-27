package notification

import (
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

func (timers Timers) GetPid(alias string) int {
	// TODO: handle case where alias does not exist
	return timers.Pids[alias]
}

func (timers *Timers) UpsertPid(alias string, pid int) {
	if timers.Pids == nil {
		timers.Pids = make(map[string]int)
	}

	timers.Pids[alias] = pid
}

func LoadTimers(pidFilePath string) Timers {
	content, _ := ioutil.ReadFile(pidFilePath)
	// TODO: handle read file error

	timers := Timers{}
	yaml.Unmarshal(content, &timers)

	return timers
}

func SaveTimers(pidFilePath string, timers Timers) {
	pidFile, _ := os.OpenFile(pidFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	defer pidFile.Close()

	content, _ := yaml.Marshal(timers)
	pidFile.Write(content)
}
