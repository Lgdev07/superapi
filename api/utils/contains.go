package utils

// Contains returns a boolean value indicating whether
// the given string contains value in the given range
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
