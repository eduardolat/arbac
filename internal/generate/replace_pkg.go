package generate

import "strings"

func replacePkg(code string, oldPkg string, newPkg string) string {
	return strings.ReplaceAll(code, "package "+oldPkg, "package "+newPkg)
}
