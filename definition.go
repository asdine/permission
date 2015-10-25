package permission

// Definition defines a Permission and its subset.
// It allows to explicitly define the rules of a permission and to test permissions against the definition.
type Definition struct {
	// Name is the name of the Permission
	Name string

	// Subset is a list of all allowed sub permissions
	Subset []string

	// DefaultSubset is a list of sub permissions allowed when only the name of the permission is specified
	DefaultSubset []string
}

// Match detects if the given permission matches the Definition
func (def *Definition) Match(perm Permission) bool {
	if perm.Name != def.Name {
		return false
	}

	if perm.Sub != "" {
		return InStringSlice(def.Subset, perm.Sub)
	}

	return true
}

// Allowed checks wether given respects required and the definition
func (def *Definition) Allowed(required, given Permission) bool {
	if required.Name != def.Name || required.Name != given.Name {
		return false
	}

	if required.Sub == "" && given.Sub == "" {
		return true
	}

	if given.Sub != "" {
		if required.Sub == "" {
			return InStringSlice(def.DefaultSubset, given.Sub)
		}
		return required.Sub == given.Sub && InStringSlice(def.Subset, given.Sub)
	}

	return InStringSlice(def.DefaultSubset, required.Sub)
}
