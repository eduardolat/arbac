package template

import "testing"

func TestPermExists(t *testing.T) {
	p := Perms{
		{
			Name: "Perm1",
			Desc: "Perm1 description",
		},
		{
			Name: "Perm2",
			Desc: "Perm2 description",
		},
		{
			Name: "Perm3",
			Desc: "Perm3 description",
		},
	}

	table := []struct {
		perm     string
		expected bool
	}{
		{
			perm:     "Perm1",
			expected: true,
		},
		{
			perm:     "Perm2",
			expected: true,
		},
		{
			perm:     "Perm3",
			expected: true,
		},
		{
			perm:     "NonExistent",
			expected: false,
		},
	}

	for _, tc := range table {
		t.Run(tc.perm, func(t *testing.T) {
			result := p.PermExistsByName(Perm{Name: tc.perm})
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
