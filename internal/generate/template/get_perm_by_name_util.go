package template

// GetPermByName returns the Perm with the given name.
func (p Perms) GetPermByName(name string) Perm {
	for _, perm := range p {
		if perm.Name == name {
			return perm
		}
	}
	return Perm{}
}
