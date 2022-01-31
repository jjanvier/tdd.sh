package handler

import (
	"errors"
	"os"
)

type ConfigHandlerI interface {
	HandleInit() error
}

type ConfigHandler struct {
	Path string
}

func (handler ConfigHandler) HandleInit() error {
	configFileExists := Configuration{}.Exists(handler.Path)
	if configFileExists {
		return errors.New("The configuration file \"" + handler.Path + "\" already exists.")
	}

	configFile, err := os.OpenFile(handler.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer configFile.Close()
	if err != nil {
		return err
	}

	_, err2 := configFile.WriteString(defaultConfigurationFile)
	if err2 != nil {
		return err2
	}

	return nil
}
