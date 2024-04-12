package schema

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed config.json
var ConfigSchema []byte

//go:embed config_template.json
var ConfigTemplate []byte

type Config struct {
	Perms   []string `json:"perms"`
	Outdir  string   `json:"outdir"`
	Package string   `json:"package"`
}

func ParseAndValidateConfig(data []byte) (Config, error) {
	if err := ValidateJSONWithSchema(ConfigSchema, data); err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, fmt.Errorf("error unmarshaling configuration: %w", err)
	}

	return config, nil
}
