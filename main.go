// Copyright (c) Div Shekhar & AUTHORS

package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "add",
				Usage: "Add a development domain and matching development port",
				Action: func(cCtx *cli.Context) error {
					return cli.Exit("Not Implemented", 1)
				},
			},
			{
				Name:   "status",
				Usage:  "Show current configuration and https server status",
				Action: status,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
