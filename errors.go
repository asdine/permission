package permission

import "errors"

// Errors
var (
	ErrEmptyName  = errors.New("The permission name is empty")
	ErrEmptyInput = errors.New("The given input is an empty string")
	ErrBadFormat  = errors.New("The given input is not in the correct format")
)
