package template

// CheckAnyPerm checks if any of the required perms
// are contained in the user perms list.
//
// Returns true if any of the required perms is found
// in the user perms list.
func (p Perms) CheckSomePerms(
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

	// Check for each required perm in the user perms list.
	for _, reqPerm := range requiredPerms {
		for _, userPerm := range userPerms {
			if reqPerm.Name == userPerm {
				return true
			}
		}
	}

	// Return false if none of the required perms are found.
	return false
}
