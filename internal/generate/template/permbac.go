package template

/*
	The permissions are statically defined in the JSON files, then
	you operate on them using the methods defined in this file.

	You can store the permissions in a database or other storage and
	associate them with users, roles or other entities as needed.

	You only need to store the permission names because the descriptions
	are only for IDE documentation purposes.
*/

// Perm represents a permission.
//
// Can be associated directly with a user or can be associated
// with a role that is then associated with a user, it's up to
// you how you want to implement it.
type Perm struct {
	Name string `json:"name" param:"name" query:"name" form:"name" xml:"name" csv:"name"`
	Desc string `json:"desc" param:"desc" query:"desc" form:"desc" xml:"desc" csv:"desc"`
}

// Perms is a slice of permissions
//
// Can be associated directly with a user or can be associated
// with a role that is then associated with a user, it's up to
// you how you want to implement it.
type Perms []Perm

// PermBAC is the instance of the PermBAC access control model.
var PermBAC = Perms{
	PermRoot,
	//*NAMES_HERE*//
}

// PermRoot Has all permissions
//
// Name: *
var PermRoot = Perm{
	Name: "*",
	Desc: "Has all permissions",
}

//*PERMS_HERE*//

// String returns a string representation of a permission.
//
// It returns the name of the permission.
func (p Perm) String() string {
	return p.Name
}

// Check checks if a permission is contained in a list of permissions.
//
// Also returns true if the root permission is found.
func (p Perm) Check(permsToCheck ...string) bool {
	for _, perm := range permsToCheck {
		if perm == p.Name || perm == PermRoot.Name {
			return true
		}
	}

	return false
}

// String returns a string representation of a list of permissions.
//
// It returns a string of permission names separated by commas.
func (p Perms) String() string {
	str := ""
	for i, perm := range p {
		if i > 0 {
			str += ", "
		}
		str += perm.Name
	}
	return str
}

// PermNames returns all permission names in a slice that can be
// iterated or stored in a database or other storage and then checked
// using the other methods.
func (p Perms) PermNames() []string {
	names := make([]string, len(p))
	for i, perm := range p {
		names[i] = perm.Name
	}
	return names
}

// CheckPerm checks if a list of user perm names contains
// a required perm.
//
// Returns true if the required perm is contained.
//
// Also returns true if the root permission is found.
func (p Perms) CheckPerm(
	permsToCheck []string,
	requiredPerm Perm,
) bool {
	// If either list is empty, return false immediately.
	if len(permsToCheck) == 0 {
		return false
	}

	// Look for root permission before the regular check.
	if PermRoot.Check(permsToCheck...) {
		return true
	}

	// Look for the required permission.
	for _, userPerm := range permsToCheck {
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
//
// Also returns true if the root permission is found.
func (p Perms) CheckAllPerms(
	permsToCheck []string,
	requiredPerms []Perm,
) bool {
	// If either list is empty, return false immediately.
	if len(permsToCheck) == 0 || len(requiredPerms) == 0 {
		return false
	}

	// Look for root permission before the regular check.
	if PermRoot.Check(permsToCheck...) {
		return true
	}

	for _, reqPerm := range requiredPerms {
		found := false
		for _, userPerm := range permsToCheck {
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
//
// Also returns true if the root permission is found.
func (p Perms) CheckAnyPerm(
	permsToCheck []string,
	requiredPerms []Perm,
) bool {
	// If either list is empty, return false immediately.
	if len(permsToCheck) == 0 || len(requiredPerms) == 0 {
		return false
	}

	// Look for root permission before the regular check.
	if PermRoot.Check(permsToCheck...) {
		return true
	}

	// Check for each required perm in the user perms list.
	for _, reqPerm := range requiredPerms {
		for _, userPerm := range permsToCheck {
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
func (p Perms) GetPermByName(name string) (Perm, bool) {
	for _, perm := range p {
		if perm.Name == name {
			return perm, true
		}
	}
	return Perm{}, false
}
