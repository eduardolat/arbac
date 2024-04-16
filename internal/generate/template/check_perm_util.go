package template

// CheckPerm checks if a list of user perm names contains
// a required perm.
//
// Returns true if the required perm is contained.
func (p Perms) CheckPerm(
	userPerms []string,
	requiredPerm Perm,
) bool {
	// If either list is empty, return false immediately.
	if len(userPerms) == 0 {
		return false
	}

	// Look for root permission before the regular check.
	for _, userPerm := range userPerms {
		if userPerm == "*" {
			return true
		}
	}

	// Look for the required permission.
	for _, userPerm := range userPerms {
		if requiredPerm.Name == userPerm {
			return true
		}
	}

	// If the required permission is not found, return false immediately.
	return false
}
