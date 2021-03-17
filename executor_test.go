package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	cmd := Command{"ls", []string{"-al", ">/dev/null"}}
	assert.Equal(t, "ls -al >/dev/null", cmd.String())
}
