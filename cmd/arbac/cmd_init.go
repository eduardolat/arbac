package main

import (
	_ "embed"
	"log"
	"os"
)

//go:embed arbac_template.json
var arbacConfig []byte

//go:embed arbac_perms_template.json
var arbacPerms []byte

const (
	configFile = "./arbac.json"
	permsFile  = "./arbac_perms.json"
)

func initCmd() {
	configFileExists, err := fileExists(configFile)
	if err != nil {
		log.Fatal(err)
	}
	if configFileExists {
		log.Fatalf("configuration file (%s) already exists", configFile)
	}

	permsFileExists, err := fileExists(permsFile)
	if err != nil {
		log.Fatal(err)
	}
	if permsFileExists {
		log.Fatalf("permissions file (%s) already exists", permsFile)
	}

	err = os.WriteFile(configFile, arbacConfig, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("created configuration file: %s", configFile)

	err = os.WriteFile(permsFile, arbacPerms, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("created permissions file: %s", permsFile)
}
