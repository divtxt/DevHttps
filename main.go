package main

import (
	"flag"
	"log"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		// TODO: show help instead
		log.Fatal("Please specify a subcommand.")
	}
}
