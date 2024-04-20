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
		"Initialize a new PermBAC configuration file",
	)
	flag.BoolVar(
		&generateFlag,
		"generate",
		false,
		"Runs the PermBAC code generator using the configuration file",
	)
	flag.StringVar(
		&configFileFlag,
		"config",
		"./permbac.json",
		"Path to the configuration file",
	)
	flag.Parse()

	fmt.Printf("\n🛡️  PermBAC %s\n", version)

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
