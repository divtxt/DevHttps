package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		// TODO: show help instead
		log.Fatal("Please specify a subcommand.")
	}
	cmd, args := args[0], args[1:]

	fmt.Printf("Command: %s\n", cmd)
}
