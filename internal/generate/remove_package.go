package generate

import "strings"

func removePackage(code string, pkg string) string {
	return strings.ReplaceAll(code, "package "+pkg, "")
}
