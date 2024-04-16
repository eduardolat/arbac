package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPermByNameFn(t *testing.T) {
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

	// Define a set of test cases
	testCases := []struct {
		name     string
		expected Perm
	}{
		{
			name:     "Perm1",
			expected: Perm{Name: "Perm1", Desc: "Perm1 description"},
		},
		{
			name:     "NonExistent",
			expected: Perm{},
		},
	}

	// Run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function and check the result
			result := p.GetPermByName(tc.name)
			assert.Equal(t, tc.expected, result)
		})
	}
}
