package main

import (
	"flag"
	"fmt"
)

const version = "v0.1.0"

var (
	initFlag       bool
	generateFlag   bool
	configFileFlag string
)

func main() {
	flag.BoolVar(
		&initFlag,
		"init",
		false,
		"Initialize a new ARBAC configuration file",
	)
	flag.BoolVar(
		&generateFlag,
		"generate",
		false,
		"Runs the ARBAC code generator using the configuration file",
	)
	flag.StringVar(
		&configFileFlag,
		"config",
		"./arbac.json",
		"Path to the configuration file",
	)
	flag.Parse()

	fmt.Printf("\n🛡️  ARBAC %s\n", version)

	if initFlag {
		initCmd()
		return
	}

	if generateFlag {
		generateCmd(configFileFlag)
		return
	}

	fmt.Println("Please provide a valid command. Run with -h for help.")
}
