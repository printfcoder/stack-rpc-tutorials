package slices

// Contains return if the 'in' slice contains the 'search' string
func Contains(in []string, search string) (b bool) {

	if in == nil || len(in) == 0 {
		return
	}

	for i := 0; i < len(in); i++ {
		if in[i] == search {
			return true
		}
	}
	return
}
