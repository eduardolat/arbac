package main

import (
	_ "embed"
	"log"
	"os"

	"github.com/eduardolat/arbac/internal/fileutil"
	"github.com/eduardolat/arbac/internal/schema"
)

const (
	configFile = "./arbac.json"
	permsFile  = "./arbac_perms.json"
)

func initCmd() {
	configFileExists, err := fileutil.FileExists(configFile)
	if err != nil {
		log.Fatal(err)
	}
	if configFileExists {
		log.Fatalf("configuration file (%s) already exists", configFile)
	}

	permsFileExists, err := fileutil.FileExists(permsFile)
	if err != nil {
		log.Fatal(err)
	}
	if permsFileExists {
		log.Fatalf("permissions file (%s) already exists", permsFile)
	}

	err = os.WriteFile(configFile, schema.ConfigTemplate, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("created configuration file: %s", configFile)

	err = os.WriteFile(permsFile, schema.PermsTemplate, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("created permissions file: %s", permsFile)
}
