package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPerm(t *testing.T) {
	tests := []struct {
		name           string
		userPerms      []string
		requiredPerm   Perm
		expectedResult bool
	}{
		{
			name:           "user has required permission",
			userPerms:      []string{"read", "write"},
			requiredPerm:   Perm{Name: "write"},
			expectedResult: true,
		},
		{
			name:           "user has root permission",
			userPerms:      []string{"*"},
			requiredPerm:   Perm{Name: "write"},
			expectedResult: true,
		},
		{
			name:           "user does not have required permission",
			userPerms:      []string{"read"},
			requiredPerm:   Perm{Name: "write"},
			expectedResult: false,
		},
		{
			name:           "required perm is empty",
			userPerms:      []string{"read", "write"},
			requiredPerm:   Perm{},
			expectedResult: false,
		},
		{
			name:           "user perms are empty",
			userPerms:      []string{},
			requiredPerm:   Perm{Name: "write"},
			expectedResult: false,
		},
		{
			name:           "both perms are empty",
			userPerms:      []string{},
			requiredPerm:   Perm{},
			expectedResult: false,
		},
	}

	p := Perms{}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := p.CheckPerm(tc.userPerms, tc.requiredPerm)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
