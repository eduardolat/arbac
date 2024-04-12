package schema

import _ "embed"

//go:embed config.json
var ConfigSchema []byte

//go:embed config_template.json
var ConfigTemplate []byte

type Config struct {
	Perms   []string `json:"perms"`
	Outdir  string   `json:"outdir"`
	Package string   `json:"package"`
}
