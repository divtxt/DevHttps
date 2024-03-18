package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

var ErrMissingConfig = errors.New("app config does not exist - please run the \"init\" command")

type AppStateDir struct {
	ConfigFile           string
	ConfigDirOrFileError error
}

func InitAppStateDir() (*AppStateDir, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	err = checkIsDir(userConfigDir)
	if err != nil {
		return nil, err
	}

	appConfigDir := filepath.Join(userConfigDir, "devhttps")
	configFile := filepath.Join(appConfigDir, "config.json")

	configErr := checkIsDir(appConfigDir)
	if configErr == nil {
		_, configErr = os.Stat(configFile)
	}
	if errors.Is(configErr, os.ErrNotExist) {
		configErr = ErrMissingConfig
	}

	return &AppStateDir{
		ConfigFile:           configFile,
		ConfigDirOrFileError: configErr,
	}, nil
}

type DevDomain struct {
	Domain    string // `json:"domain"`
	LocalPort int    // `json:"LocalPort"`
}

type AppConfig struct {
	Version string      //`json:"version""`
	Domains []DevDomain //`json:"domains"`
}

func (asd *AppStateDir) LoadConfig() (*AppConfig, error) {
	if asd.ConfigDirOrFileError != nil {
		return nil, asd.ConfigDirOrFileError
	}

	f, err := os.Open(asd.ConfigFile)
	if err != nil {
		return nil, err
	}

	var appConfig AppConfig
	jsonParser := json.NewDecoder(f)
	if err = jsonParser.Decode(&appConfig); err != nil {
		return nil, err
	}

	return &appConfig, nil
}

func checkIsDir(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !fi.Mode().IsDir() {
		return errors.New("\"" + path + "\" is not a directory")
	}
	return nil
}
