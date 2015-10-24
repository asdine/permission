# Permission

Permission is a low-level Go package that allows to easily manage permissions.

## Install

```
$ go get -u github.com/asdine/permission
```

## Permission

Permission is the primitive that defines a single permission.

```go
p := permission.New("read")
```

You can specify a sub Permission by adding a `.`

```go
p := permission.New("user.edit")
```

The `.` delimiter can be changed by setting the global package delimiter

```go
permission.Delimiter(":")

p := permission.New("user:edit")
```

The variable return by `permission.New` is a Permission primitive than can be easily manipulated and marshalled.

```go
p := permission.New("user.edit")
q := permission.New("user.edit")

fmt.Println(p.Name)
// user

fmt.Println(p.Sub)
// edit

fmt.Println(p.Equal(q))
// true

fmt.Println(p)
// user.edit
```

The primitive can be Unmarshalled from JSON ...

```go
type Access struct {
	Name       string
	Permission permission.Permission
}

a := Access{}

input := []byte(`{"Name":"Edit User","Permission":"user.edit"}`)
json.Unmarshal(input, &a)

fmt.Println(a.Permission.Name)
// user
fmt.Println(a.Permission.Sub)
// edit
```

... and marshalled to JSON

```go
output, _ := json.Marshal(a)
fmt.Println(output)
// {"Name":"Edit User","Permission":"user.edit"}
```
