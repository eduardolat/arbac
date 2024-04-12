package schema

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed perms.json
var PermsSchema []byte

//go:embed perms_template.json
var PermsTemplate []byte

type Perm struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Perms []Perm

func ParseAndValidatePerms(dataSlice [][]byte) (Perms, error) {
	allPerms := Perms{}

	for _, data := range dataSlice {
		if err := ValidateJSONWithSchema(data, PermsSchema); err != nil {
			return Perms{}, err
		}

		var perms Perms
		if err := json.Unmarshal(data, &perms); err != nil {
			return Perms{}, fmt.Errorf("error unmarshaling permissions: %w", err)
		}

		allPerms = append(allPerms, perms...)
	}

	if len(allPerms) == 0 {
		return Perms{}, fmt.Errorf("no permissions found")
	}

	permsNames := map[string]bool{}
	for _, perm := range allPerms {
		if _, ok := permsNames[perm.Name]; ok {
			return Perms{}, fmt.Errorf("duplicate permission name: %s", perm.Name)
		}
		permsNames[perm.Name] = true
	}

	return allPerms, nil
}
