package main

import (
	"flag"
	"fmt"
)

const version = "v1.0.0"

var (
	versionFlag    bool
	initFlag       bool
	generateFlag   bool
	configFileFlag string
)

func main() {
	flag.BoolVar(
		&versionFlag,
		"version",
		false,
		"Prints the version of PermBAC",
	)
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

	fmt.Printf("🛡️  PermBAC %s\n", version)
	if versionFlag {
		return
	}

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
