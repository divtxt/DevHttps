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
				Name:  "status",
				Usage: "Show server configuration and current status",
				Action: func(cCtx *cli.Context) error {
					return cli.Exit("Not Implemented", 1)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
