package permission

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalling(t *testing.T) {
	p := Permission{Name: "a"}

	output, err := p.MarshalText()
	assert.NoError(t, err)
	assert.Equal(t, "a", string(output))

	p = Permission{Name: "a", Sub: "b"}

	output, err = p.MarshalText()
	assert.NoError(t, err)
	assert.Equal(t, "a.b", string(output))

	p = Permission{}

	output, err = p.MarshalText()
	assert.Error(t, err)
	assert.Equal(t, ErrEmptyName, err)
	assert.Nil(t, output)
}

func TestUnMarshalling(t *testing.T) {
	text := []byte("a")
	p := Permission{}

	err := p.UnmarshalText(text)
	assert.NoError(t, err)
	assert.Equal(t, "a", p.Name)
	assert.Empty(t, p.Sub)

	text = []byte("a.b")
	p = Permission{}

	err = p.UnmarshalText(text)
	assert.NoError(t, err)
	assert.Equal(t, "a", p.Name)
	assert.Equal(t, "b", p.Sub)

	err = p.UnmarshalText(nil)
	assert.Error(t, err)
	assert.Equal(t, ErrEmptyInput, err)

	text = []byte("a.")
	err = p.UnmarshalText(text)
	assert.Error(t, err)
	assert.Equal(t, ErrBadFormat, err)

	text = []byte(".b")
	err = p.UnmarshalText(text)
	assert.Error(t, err)
	assert.Equal(t, ErrBadFormat, err)
}

func TestPermToJSON(t *testing.T) {
	p := Permission{Name: "a"}

	val, err := json.Marshal(p)
	assert.NoError(t, err)
	assert.Equal(t, `"a"`, string(val))

	p = Permission{Name: "a", Sub: "b"}

	val, err = json.Marshal(p)
	assert.NoError(t, err)
	assert.Equal(t, `"a.b"`, string(val))

	p = Permission{}
	val, err = json.Marshal(p)
	assert.Error(t, err)
}

func TestPermFromJSON(t *testing.T) {
	s := []byte(`"a"`)
	p := Permission{}

	err := json.Unmarshal(s, &p)
	assert.NoError(t, err)
	assert.Equal(t, "a", p.Name)

	s = []byte(`"a.b"`)
	p = Permission{}

	err = json.Unmarshal(s, &p)
	assert.NoError(t, err)
	assert.Equal(t, "a", p.Name)
	assert.Equal(t, "b", p.Sub)

	s = []byte(`""`)
	p = Permission{}

	err = json.Unmarshal(s, &p)
	assert.Error(t, err)
}

func TestPermToFromJSON(t *testing.T) {
	s := Permission{Name: "a", Sub: "b"}

	val, err := json.Marshal(s)
	assert.NoError(t, err)

	d := Permission{}
	err = json.Unmarshal(val, &d)
	assert.NoError(t, err)
	assert.Equal(t, s, d)
}

func Example() {
	permission := Permission{
		Name: "user",
		Sub:  "edit",
	}

	data, _ := json.Marshal(permission)
	fmt.Printf("%s\n", data)

	p := Permission{}
	json.Unmarshal(data, &p)
	fmt.Println(p.Name)
	fmt.Println(p.Sub)
	// Output:
	// "user.edit"
	// user
	// edit
}

func ExamplePermission() {
	// Simple Permission
	perm := Permission{Name: "read"}
	fmt.Println(perm)
	// Output: read
}

func ExamplePermission_subPermission() {
	// Sub Permission
	perm := Permission{Name: "user", Sub: "edit"}
	fmt.Println(perm)
	// Output: user.edit
}
