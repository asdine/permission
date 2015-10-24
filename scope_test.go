package permission

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScopeMarshalling(t *testing.T) {
	s := Scope{
		Permission{Name: "a"},
		Permission{Name: "b", Sub: "i"},
		Permission{Name: "c", Sub: "j"},
	}

	output, err := s.MarshalText()
	assert.NoError(t, err)
	assert.Equal(t, "a,b.i,c.j", string(output))

	s = Scope{}
	output, err = s.MarshalText()
	assert.NoError(t, err)

	s = Scope{Permission{}}
	output, err = s.MarshalText()
	assert.Error(t, err)
	assert.Equal(t, ErrEmptyName, err)
}

func TestScopeUnmarshalling(t *testing.T) {
	input := []byte("a,b.i,c.j, d")
	expected := Scope{
		Permission{Name: "a"},
		Permission{Name: "b", Sub: "i"},
		Permission{Name: "c", Sub: "j"},
		Permission{Name: " d"},
	}

	scope := Scope{}
	err := scope.UnmarshalText(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, scope)

	input = []byte("")
	scope = Scope{}
	err = scope.UnmarshalText(input)
	assert.Error(t, err)
	assert.Equal(t, ErrEmptyInput, err)

	input = []byte("a,b.")
	scope = Scope{}
	err = scope.UnmarshalText(input)
	assert.Error(t, err)
	assert.Equal(t, ErrBadFormat, err)
}

func TestScopeToJSON(t *testing.T) {
	s := Scope{
		Permission{Name: "a"},
		Permission{Name: "b", Sub: "i"},
		Permission{Name: "c", Sub: "j"},
	}

	val, err := json.Marshal(s)
	assert.NoError(t, err)
	assert.Equal(t, `"a,b.i,c.j"`, string(val))
}

func TestScopeFromJSON(t *testing.T) {
	input := []byte(`"a,b.i,c.j, d, e.k"`)
	expected := Scope{
		Permission{Name: "a"},
		Permission{Name: "b", Sub: "i"},
		Permission{Name: "c", Sub: "j"},
		Permission{Name: " d"},
		Permission{Name: " e", Sub: "k"},
	}

	s := Scope{}
	err := json.Unmarshal(input, &s)
	assert.NoError(t, err)
	assert.Equal(t, expected, s)
}

func TestSepatator(t *testing.T) {
	Separator(":")
	defer Separator(",")

	s := Scope{
		Permission{Name: "a"},
		Permission{Name: "b", Sub: "i"},
	}

	output, err := s.MarshalText()
	assert.NoError(t, err)
	assert.Equal(t, "a:b.i", string(output))
}
