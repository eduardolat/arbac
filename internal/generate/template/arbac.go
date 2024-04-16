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

//*UTILS_HERE*//
