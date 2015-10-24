// Package permission is a low-level Go package that allows to easily manage permissions
package permission

import (
	"fmt"
	"strings"
)

// Permission is a simple permission structure.
// It is meant to describe a single specific permission.
// A Sub permission can be specified.
// It can safely be converted back and forth to json
type Permission struct {
	Name string
	Sub  string
}

// MarshalText encodes the given text
func (p Permission) MarshalText() (text []byte, err error) {
	if p.Name == "" {
		return nil, ErrEmptyName
	}

	if p.Sub == "" {
		return []byte(p.Name), nil
	}

	return []byte(fmt.Sprintf("%s.%s", p.Name, p.Sub)), nil
}

// UnmarshalText decodes the given text
func (p *Permission) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		return ErrEmptyInput
	}

	perm := string(text)
	if !strings.Contains(perm, ".") {
		p.Name = perm
		return nil
	}

	frags := strings.Split(perm, ".")
	if len(frags) != 2 {
		return ErrBadFormat
	}

	if frags[0] == "" || frags[1] == "" {
		return ErrBadFormat
	}

	p.Name = frags[0]
	p.Sub = frags[1]
	return nil
}

// String returns the string representation of the Permission
func (p Permission) String() string {
	if p.Sub != "" {
		return fmt.Sprintf("%s.%s", p.Name, p.Sub)
	}

	return p.Name
}
