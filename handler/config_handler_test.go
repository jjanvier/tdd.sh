package handler

import (
	"github.com/jjanvier/tdd/helper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestHandleInit(t *testing.T) {
	file := helper.TmpFileName()
	defer helper.RemoveTmpFileByName(file)

	handler := ConfigHandler{file}
	handler.HandleInit()

	content, _ := ioutil.ReadFile(file)

	assert.Equal(t, defaultConfigurationFile, string(content))
}

func TestDoesNotWipeOutFileDuringInit(t *testing.T) {
	file := helper.CreateTmpFile("my already existing configuration file")
	defer helper.RemoveTmpFile(file)

	handler := ConfigHandler{file.Name()}
	err := handler.HandleInit()

	assert.Error(t, err)
}

func TestHandleValidate(t *testing.T) {
	file := helper.CreateTmpFile(defaultConfigurationFile)
	defer helper.RemoveTmpFile(file)

	handler := ConfigHandler{file.Name()}

	assert.True(t, handler.HandleValidate())
}

func TestHandleNotValidate(t *testing.T) {
	file := helper.CreateTmpFile("wrong configuration here!")
	defer helper.RemoveTmpFile(file)

	handler := ConfigHandler{file.Name()}

	assert.False(t, handler.HandleValidate())
}
