package permission

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchSimpleDef(t *testing.T) {
	d := Definition{Name: "a"}

	p := Permission{Name: "b"}
	assert.False(t, d.Match(p))

	p = Permission{Name: ""}
	assert.False(t, d.Match(p))

	p = Permission{Name: "a"}
	assert.True(t, d.Match(p))
}

func TestMatchDef(t *testing.T) {
	d := Definition{
		Name:          "a",
		Subset:        []string{"i", "j", "k"},
		DefaultSubset: []string{"i", "j"},
	}

	p := Permission{Name: "a"}
	assert.True(t, d.Match(p))

	p = Permission{Name: "b"}
	assert.False(t, d.Match(p))

	p = Permission{Name: ""}
	assert.False(t, d.Match(p))

	p = Permission{Name: "a", Sub: "l"}
	assert.False(t, d.Match(p))

	p = Permission{Name: "b", Sub: "i"}
	assert.False(t, d.Match(p))

	p = Permission{Name: "a", Sub: "i"}
	assert.True(t, d.Match(p))

	p = Permission{Name: "a", Sub: "j"}
	assert.True(t, d.Match(p))

	p = Permission{Name: "a", Sub: "k"}
	assert.True(t, d.Match(p))
}

func TestAllowedSimple(t *testing.T) {
	d := Definition{
		Name:          "a",
		Subset:        []string{"i", "j", "k"},
		DefaultSubset: []string{"i", "j"},
	}

	required := Permission{Name: "a"}

	p := Permission{Name: "b"}
	assert.False(t, d.Allowed(required, p))

	p = Permission{Name: ""}
	assert.False(t, d.Allowed(required, p))

	p = Permission{Name: "a"}
	assert.True(t, d.Allowed(required, p))

	p = Permission{Name: "a", Sub: "i"}
	assert.True(t, d.Allowed(required, p))

	required = Permission{Name: ""}
	p = Permission{Name: "a"}
	assert.False(t, d.Allowed(required, p))
}

func TestAllowed(t *testing.T) {
	d := Definition{
		Name:          "a",
		Subset:        []string{"i", "j", "k"},
		DefaultSubset: []string{"i", "j"},
	}

	required := Permission{Name: "a", Sub: "i"}

	p := Permission{Name: "a", Sub: "j"}
	assert.False(t, d.Allowed(required, p))

	p = Permission{Name: "a", Sub: "i"}
	assert.True(t, d.Allowed(required, p))

	p = Permission{Name: "a"} // a = a.i,a.j
	assert.True(t, d.Allowed(required, p))

	required = Permission{Name: "a", Sub: "k"}

	p = Permission{Name: "a"}
	assert.False(t, d.Allowed(required, p))
}
