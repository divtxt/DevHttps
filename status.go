package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func status(cCtx *cli.Context) error {
	appStateDir, err := InitAppStateDir()
	if err != nil {
		return cli.Exit(err, 1)
	}

	// Configuration file
	fmt.Printf("ConfigFile: \"%s\"\n", appStateDir.ConfigFile)
	if appStateDir.ConfigDirOrFileError != nil {
		return cli.Exit(appStateDir.ConfigDirOrFileError, 1)
	}

	// Configuration
	appConfig, err := appStateDir.LoadConfig()
	if err != nil {
		return cli.Exit(err, 1)
	}
	fmt.Printf("Configuration: %v\n", appConfig)

	// Server Status:
	// FIXME

	return nil
}
