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

func TestNewScope(t *testing.T) {
	s, err := ParseScope("a.b")
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Len(t, s, 1)
	assert.Equal(t, Permission{Name: "a", Sub: "b"}, s[0])
	val, err := json.Marshal(s)
	assert.NoError(t, err)
	assert.Equal(t, `"a.b"`, string(val))

	s, err = ParseScope("a.b,c")
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Len(t, s, 2)
	assert.True(t, s[1].Equal(Permission{Name: "c"}))
	val, err = json.Marshal(s)
	assert.NoError(t, err)
	assert.Equal(t, `"a.b,c"`, string(val))

	Separator(":")
	defer Separator(",")

	s, err = ParseScope("a,b")
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Len(t, s, 1)
	assert.Equal(t, "a,b", s[0].Name)

	s, err = ParseScope("a:b")
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Len(t, s, 2)
	assert.Equal(t, "a", s[0].Name)
	assert.Equal(t, "b", s[1].Name)

	s, err = ParseScope("a:")
	assert.Error(t, err)
	assert.Nil(t, s)

	Separator(".")
	Delimiter(".")

	s, err = ParseScope("a,b.c.c.e")
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Len(t, s, 4)
	assert.Equal(t, "a,b", s[0].Name)
	assert.Equal(t, "c", s[1].Name)
	assert.Equal(t, "c", s[2].Name)
	assert.Equal(t, "e", s[3].Name)
}
