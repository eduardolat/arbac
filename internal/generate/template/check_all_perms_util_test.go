package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckAllPerms(t *testing.T) {
	tests := []struct {
		name           string
		userPerms      []string
		requiredPerms  []Perm
		expectedResult bool
	}{
		{
			name:           "user has required permission",
			userPerms:      []string{"read", "write"},
			requiredPerms:  []Perm{{Name: "write"}},
			expectedResult: true,
		},
		{
			name:           "user has multiple required permission",
			userPerms:      []string{"read", "write", "delete"},
			requiredPerms:  []Perm{{Name: "write"}, {Name: "delete"}},
			expectedResult: true,
		},
		{
			name:           "user has root permission",
			userPerms:      []string{"*"},
			requiredPerms:  []Perm{{Name: "write"}, {Name: "delete"}},
			expectedResult: true,
		},
		{
			name:           "user does not have required permission",
			userPerms:      []string{"read"},
			requiredPerms:  []Perm{{Name: "write"}},
			expectedResult: false,
		},
		{
			name:           "user only have 1 of 2 required permission",
			userPerms:      []string{"read", "write"},
			requiredPerms:  []Perm{{Name: "write"}, {Name: "delete"}},
			expectedResult: false,
		},
		{
			name:           "required perms are empty",
			userPerms:      []string{"read", "write"},
			requiredPerms:  []Perm{},
			expectedResult: false,
		},
		{
			name:           "user perms are empty",
			userPerms:      []string{},
			requiredPerms:  []Perm{{Name: "write"}},
			expectedResult: false,
		},
		{
			name:           "both perms are empty",
			userPerms:      []string{},
			requiredPerms:  []Perm{},
			expectedResult: false,
		},
	}

	p := Perms{}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := p.CheckAllPerms(tc.userPerms, tc.requiredPerms)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
