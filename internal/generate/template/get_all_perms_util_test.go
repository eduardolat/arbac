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
