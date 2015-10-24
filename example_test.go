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

	permission.Delimiter(":")
	perm = permission.Permission{Name: "user", Sub: "edit"}
	fmt.Println(perm)
	// Output:
	// user.edit
	// user:edit
}
