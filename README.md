# Permission

[![Build Status](https://travis-ci.org/asdine/permission.svg)](https://travis-ci.org/asdine/permission)
[![GoDoc](https://godoc.org/github.com/asdine/permission?status.svg)](https://godoc.org/github.com/asdine/permission)
![License MIT](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)

Permission is a low-level Go package that allows to easily manage permissions.

## Install

```
$ go get -u github.com/asdine/permission
```

## Permission

Permission is the primitive that defines a single permission.

```go
p, _ := permission.Parse("read")
```

You can specify a sub Permission by adding a `.`

```go
p, _ := permission.Parse("user.edit")
```

The `.` delimiter can be changed by setting the global package delimiter

```go
permission.Delimiter(":")

p, _ := permission.Parse("user:edit")
```

The variable returned by `permission.Parse` is a Permission primitive than can be easily manipulated and marshalled.

```go
p, _ := permission.Parse("user.edit")
q, _ := permission.Parse("user.edit")

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

## Scope

A Scope is a set of permissions. It can be used to describe multiple permissions.

```go
s, _ := permission.ParseScope("read,write,edit,user:email")
```

The `,` separator can be changed by setting the global package separator

```go
permission.Separator(" ")

s, _ := permission.ParseScope("read write edit user:email")
```

The variable returned by `permission.ParseScope` is a Scope primitive helper to manipulate sets of Permissions.

```go
s, _ := permission.ParseScope("read,write,user:email")

fmt.Println(len(s))
// 3

fmt.Println(s[0].Name)
// read

fmt.Println(s[2].Sub)
// email

fmt.Println(s)
// read,write,edit,user:email
```

JSON example
```go
type Role struct {
  Name        string
  Permissions permission.Scope
}

r := Role{}

input := []byte(`{"Name":"Admin","Permission":"read,write,user:email"}`)
json.Unmarshal(input, &r)

fmt.Println(len(r.Permissions))
// 3

fmt.Println(r.Permissions[0].Name)
// read

fmt.Println(a.Permissions[2].Sub)
// edit

output, _ := json.Marshal(r)
fmt.Println(output)
// {"Name":"Admin","Permission":"read,write,user:email"}
```

## Definition

Definition is a way of defining permission attributes and rules.

```go
def := permission.Definition{
	Name:          "playlist",
	Subset:        []string{"edit", "share", "read"},
	DefaultSubset: []string{"read", "share"},
}
```

Once a definition is created you can test permissions against it to see if they match

```go
p, _ := permission.Parse("playlist.edit")

fmt.Println(def.Match(p))
// true
```

You can also compare two permissions and test them against the definition.
It is useful when you have a permission that has default sub permissions.

```go
required, _ := permission.Parse("playlist")
// required gets granted the DefaultSubset list of permissions.
// It is equivalent to playlist.read and playlist.share

p, _ := permission.Parse("playlist.share")

fmt.Println(def.Allowed(required, p))
// true
```

The given permission can also benefits of the DefaultSubset

```go
required, _ := permission.Parse("playlist.read")

p, _ := permission.Parse("playlist")
// required gets granted the DefaultSubset list of permissions.
// It is equivalent to playlist.read and playlist.share

fmt.Println(def.Allowed(required, p))
// true
```

## License

MIT

## Author

**Asdin El Hrychy**

- [Twitter](https://twitter.com/asdine_)
- [Github](https://github.com/asdine)
