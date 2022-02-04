package parser

func compareRuneSlice(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func toStringSlice(a []rune) []string {
	var result []string
	for i := range a {
		result = append(result, string(a[i]))
	}
	return result
}
