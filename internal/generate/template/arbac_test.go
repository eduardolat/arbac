package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllPerms(t *testing.T) {
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

	got := p.GetAllPerms()
	if len(got) != 3 {
		t.Errorf("Perms.GetAllPerms() = %v; want 3", len(got))
	}

	for i, perm := range got {
		assert.Equal(t, p[i].Name, perm.Name)
		assert.Equal(t, p[i].Desc, perm.Desc)
	}
}
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

func TestCheckSomePerms(t *testing.T) {
	tests := []struct {
		name           string
		userPerms      []string
		requiredPerms  []Perm
		expectedResult bool
	}{
		{
			name:           "user has required permission",
			userPerms:      []string{"read", "write"},
			requiredPerms:  []Perm{{Name: "write"}, {Name: "other1"}, {Name: "other2"}},
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
			expectedResult: true,
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
			result := p.CheckSomePerms(tc.userPerms, tc.requiredPerms)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestGetPermByName(t *testing.T) {
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
		name          string
		expected      Perm
		expectedFound bool
	}{
		{
			name:          "Perm1",
			expected:      Perm{Name: "Perm1", Desc: "Perm1 description"},
			expectedFound: true,
		},
		{
			name:          "NonExistent",
			expected:      Perm{},
			expectedFound: false,
		},
	}

	// Run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function and check the result
			found, result := p.GetPermByName(tc.name)
			assert.Equal(t, tc.expected, result)
			assert.Equal(t, tc.expectedFound, found)
		})
	}
}
