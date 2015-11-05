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

func TestDefinitions_Definition(t *testing.T) {
	d := Definitions{
		Definition{
			Name:          "a",
			Subset:        []string{"i", "j", "k"},
			DefaultSubset: []string{"i", "j"},
		},
		Definition{
			Name:          "b",
			Subset:        []string{"i", "j", "k"},
			DefaultSubset: []string{"i", "j"},
		},
	}

	assert.NotNil(t, d.Definition(Permission{Name: "a"}))
	assert.NotNil(t, d.Definition(Permission{Name: "b"}))
	assert.NotNil(t, d.Definition(Permission{Name: "a", Sub: "i"}))
	assert.Nil(t, d.Definition(Permission{Name: "a", Sub: "z"}))
	assert.Nil(t, d.Definition(Permission{Name: "c"}))
	assert.Nil(t, d.Definition(Permission{Name: ""}))
}

func TestDefinitions_Require(t *testing.T) {
	d := Definitions{
		{
			Name:          "a",
			Subset:        []string{"i", "j", "k"},
			DefaultSubset: []string{"i", "j"},
		},
		{
			Name:          "b",
			Subset:        []string{"i", "j", "k"},
			DefaultSubset: []string{"i", "j"},
		},
	}

	assert.False(t, d.Require("a", "b"))
	assert.True(t, d.Require("a", "a"))
	assert.True(t, d.Require("b", "b"))
	assert.True(t, d.Require("a,b", "b.i"))
	assert.True(t, d.Require("a,b", "a.i"))
	assert.False(t, d.Require("a,b", "a.k"))
	assert.True(t, d.Require("a,b", "a.k,b.i"))
	assert.False(t, d.Require("a.i", "a.k,b.i"))
	assert.True(t, d.Require("a.i", "a,b.i"))
	assert.False(t, d.Require("a.", "a,b.i"))
	assert.False(t, d.Require("a", "a,"))
}
