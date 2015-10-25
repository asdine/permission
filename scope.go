package permission

import (
	"bytes"
	"strings"
	"sync"
)

var separator = ","

var sepLock sync.Mutex

// Separator is a thread-safe function that sets a global separator for a set of Permissions.
// Defaults to ","
func Separator(sep string) {
	sepLock.Lock()
	defer sepLock.Unlock()
	separator = sep
}

// ParseScope takes a string representation and returns the corresponding Scope
func ParseScope(repr string) (Scope, error) {
	s := Scope{}
	err := s.UnmarshalText([]byte(repr))
	if err != nil {
		return nil, err
	}
	return s, nil
}

// Scope is a set of Permissions.
// It can safely be converted back and forth to json
type Scope []Permission

// MarshalText implements the encoding.TextMarshaler interface
func (s Scope) MarshalText() (text []byte, err error) {
	var buffer bytes.Buffer

	for i, perm := range s {
		raw, err := perm.MarshalText()
		if err != nil {
			return nil, err
		}

		buffer.Write(raw)
		if i < len(s)-1 {
			buffer.WriteString(separator)
		}
	}

	return buffer.Bytes(), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface
func (s *Scope) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		return ErrEmptyInput
	}

	scope := strings.Split(string(text), separator)
	*s = make([]Permission, len(scope))
	for i, perm := range scope {
		err := (*s)[i].UnmarshalText([]byte(perm))
		if err != nil {
			return err
		}
	}
	return nil
}
