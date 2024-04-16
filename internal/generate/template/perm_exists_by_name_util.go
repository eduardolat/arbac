package template

// PermExistsByName returns true if the given Perm exists.
func (p Perms) PermExistsByName(perm Perm) bool {
	for _, p := range p {
		if p.Name == perm.Name {
			return true
		}
	}
	return false
}
