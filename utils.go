package permission

// InStringSlice checks if the given string is in the given slice of string
func InStringSlice(haystack []string, needle string) bool {
	for _, str := range haystack {
		if needle == str {
			return true
		}
	}

	return false
}
