package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eduardolat/arbac/internal/fileutil"
	"github.com/eduardolat/arbac/internal/schema"
)

func generateCmd(configFile string) {
	log.Printf("Using %s config file\n", configFile)

	configExists, err := fileutil.FileExists(configFile)
	if err != nil {
		log.Fatalf("error checking if config file exists: %v", err)
	}
	if !configExists {
		log.Fatalf("config file %s does not exist", configFile)
	}
	configData, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	config, err := schema.ParseAndValidateConfig(configData)
	if err != nil {
		log.Fatal(err.Error())
	}

	files, err := fileutil.ReadGlobFiles(config.Perms)
	if err != nil {
		log.Fatalf("error reading permission files: %v", err)
	}

	perms, err := schema.ParseAndValidatePerms(files)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(perms)
}
