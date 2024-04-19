package template

// Perm represents a permission that can be assigned to a role or user
type Perm struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// Perms is a slice of permissions
type Perms []Perm

// Arbac is the instance of the ARBAC access control model.
var Arbac = Perms{
	PermRoot,
	//*NAMES_HERE*//
}

// PermRoot Super user, has all permissions
var PermRoot = Perm{
	Name: "*",
	Desc: "Super user, has all permissions",
}

//*PERMS_HERE*//

// GetAllPerms returns all permissions in the ARBAC instance
func (p Perms) GetAllPerms() []Perm {
	return p
}

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

// GetPermByName returns the Perm with the given name
// and a boolean indicating if it was found.
func (p Perms) GetPermByName(name string) (bool, Perm) {
	for _, perm := range p {
		if perm.Name == name {
			return true, perm
		}
	}
	return false, Perm{}
}
