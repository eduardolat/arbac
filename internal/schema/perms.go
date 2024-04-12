package schema

import _ "embed"

//go:embed perms.json
var PermsSchema []byte

//go:embed perms_template.json
var PermsTemplate []byte

type Perm struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Perms []Perm
