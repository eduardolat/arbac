package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermString(t *testing.T) {
	p := Perm{
		Name: "Perm1",
		Desc: "Perm1 description",
	}

	got := p.String()
	assert.Equal(t, p.Name, got)
}

func TestPermCheck(t *testing.T) {
	p := Perm{
		Name: "Perm1",
	}

	tests := []struct {
		permNames      []string
		expectedResult bool
	}{
		{
			permNames:      []string{"Perm1", "Perm2", "Perm3"},
			expectedResult: true,
		},
		{
			permNames:      []string{"Perm2", "Perm3"},
			expectedResult: false,
		},
		{
			permNames:      []string{PermRoot.Name},
			expectedResult: true,
		},
	}

	for _, tc := range tests {
		t.Run("MatchSomePerm", func(t *testing.T) {
			result := p.Check(tc.permNames...)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestPermsString(t *testing.T) {
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

	got := p.String()
	assert.Equal(t, "Perm1, Perm2, Perm3", got)
}

func TestPermNames(t *testing.T) {
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

	got := p.PermNames()
	want := []string{"Perm1", "Perm2", "Perm3"}

	if len(got) != 3 {
		t.Errorf("Perms.PermNames() = %v; want 3", len(got))
	}
	assert.Equal(t, want, got)
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
			userPerms:      []string{PermRoot.Name},
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
			userPerms:      []string{PermRoot.Name},
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

func TestCheckAnyPerm(t *testing.T) {
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
			userPerms:      []string{PermRoot.Name},
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
			result := p.CheckAnyPerm(tc.userPerms, tc.requiredPerms)
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
			result, found := p.GetPermByName(tc.name)
			assert.Equal(t, tc.expected, result)
			assert.Equal(t, tc.expectedFound, found)
		})
	}
}
