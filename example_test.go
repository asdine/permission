package permission_test

import (
	"encoding/json"
	"fmt"

	"github.com/asdine/permission"
)

func Example() {
	p := permission.Permission{
		Name: "user",
		Sub:  "edit",
	}

	data, _ := json.Marshal(p)
	fmt.Printf("%s\n", data)

	p = permission.Permission{}
	json.Unmarshal(data, &p)
	fmt.Println(p.Name)
	fmt.Println(p.Sub)

	def := permission.Definition{
		Name:          "user",
		Subset:        []string{"edit", "profile", "email", "friends", "about"},
		DefaultSubset: []string{"profile", "about"},
	}

	required, _ := permission.Parse("user.edit")
	allowed := def.Allowed(required, p)
	fmt.Println(allowed)
	// Output:
	// "user.edit"
	// user
	// edit
	// true
}

func ExamplePermission() {
	// Simple Permission
	perm := permission.Permission{Name: "read"}
	fmt.Println(perm)
	// Output: read
}

func ExamplePermission_subPermission() {
	// Sub Permission
	perm := permission.Permission{Name: "user", Sub: "edit"}
	fmt.Println(perm)
	// Output: user.edit
}

func ExampleParse() {
	perm, _ := permission.Parse("user.edit")

	fmt.Println(perm.Name)
	fmt.Println(perm.Sub)
	fmt.Println(perm)
	// Output:
	// user
	// edit
	// user.edit
}

func ExampleDelimiter() {
	permission.Delimiter(":")
	defer permission.Delimiter(".")
	perm := permission.Permission{Name: "user", Sub: "edit"}
	fmt.Println(perm)
	// Output: user:edit
}

func ExamplePermission_Equal() {
	p, _ := permission.Parse("user.edit")
	q, _ := permission.Parse("user.edit")

	fmt.Println(p.Equal(q))
	// Output: true
}

func ExampleScope() {
	perms := permission.Scope{
		permission.Permission{Name: "user", Sub: "edit"},
		permission.Permission{Name: "profile"},
		permission.Permission{Name: "friends"},
	}

	text, _ := perms.MarshalText()
	fmt.Println(string(text))
	// Output: user.edit,profile,friends
}

func ExampleParseScope() {
	perms, _ := permission.ParseScope("user.edit,profile,friends")

	fmt.Println(len(perms))
	fmt.Println(perms[0].Name)
	fmt.Println(perms[0].Sub)
	fmt.Println(perms[1].Name)
	fmt.Println(perms[2].Name)
	// Output:
	// 3
	// user
	// edit
	// profile
	// friends
}

func ExampleDefinition() {
	def := permission.Definition{
		Name:          "file",
		Subset:        []string{"read", "write", "execute"},
		DefaultSubset: []string{"read", "execute"},
	}

	required, _ := permission.Parse("file.read")

	p, _ := permission.Parse("file.read")
	fmt.Println(def.Allowed(required, p))

	p, _ = permission.Parse("file.execute")
	fmt.Println(def.Allowed(required, p))

	p, _ = permission.Parse("file") // file = file.read,file.execute
	fmt.Println(def.Allowed(required, p))

	required, _ = permission.Parse("file.write")

	p, _ = permission.Parse("file") // file = file.read,file.execute
	fmt.Println(def.Allowed(required, p))
	// Output:
	// true
	// false
	// true
	// false
}

func ExampleDefinition_Match() {
	defs := []permission.Definition{
		permission.Definition{
			Name:          "repo",
			Subset:        []string{"read", "write"},
			DefaultSubset: []string{"read"},
		},
		permission.Definition{
			Name:          "user",
			Subset:        []string{"profile", "edit", "friends", "email"},
			DefaultSubset: []string{"profile", "friends"},
		},
		permission.Definition{
			Name:          "playlist",
			Subset:        []string{"edit", "share", "read"},
			DefaultSubset: []string{"read", "share"},
		},
	}

	p, _ := permission.Parse("user.edit")

	for _, def := range defs {
		if def.Match(p) {
			fmt.Println(def.Name)
		}
	}
	// Output:
	// user
}
