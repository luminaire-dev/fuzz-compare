package compare

// Equal compares two slices of bytes.
func Equal(a, b []byte) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
