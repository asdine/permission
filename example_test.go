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
	// Output:
	// "user.edit"
	// user
	// edit
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

func ExampleNew() {
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

func ExampleNewScope() {
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
