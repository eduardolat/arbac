package template

// CheckAllPerms checks if a list of user perm names are
// contained in another list of required perms.
//
// Returns true if all required perms are contained
// in the user perms list.
func (p Perms) CheckAllPerms(
	userPerms []string,
	requiredPerms []Perm,
) bool {
	// If either list is empty, return false immediately.
	if len(userPerms) == 0 || len(requiredPerms) == 0 {
		return false
	}

	// Look for root permission before the regular check.
	for _, userPerm := range userPerms {
		if userPerm == "*" {
			return true
		}
	}

	for _, reqPerm := range requiredPerms {
		found := false
		for _, userPerm := range userPerms {
			if reqPerm.Name == userPerm {
				found = true
				break
			}
		}

		// If a required perm is not found, return false immediately.
		if !found {
			return false
		}
	}

	return true
}
