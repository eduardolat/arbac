package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/eduardolat/permbac/internal/fileutil"
	"github.com/eduardolat/permbac/internal/generate"
	"github.com/eduardolat/permbac/internal/schema"
)

func generateCmd(configFile string) {
	startDate := time.Now()
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

	pbytes, err := generate.GeneratePerms(version, config, perms)
	if err != nil {
		log.Fatal(err.Error())
	}

	permsPath := filepath.Join(config.Outdir, "permbac_generated.go")
	err = os.WriteFile(permsPath, pbytes, 0777)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf(
		"Generated %d permissions in %s\n",
		len(perms),
		time.Since(startDate).Round(time.Millisecond),
	)
}
