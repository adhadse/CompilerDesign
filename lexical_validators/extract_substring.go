package lexical_validators

// SubString extracts the SUBSTRING from position
// left to right(excluding) index.
func SubString(s string, left, right int) string {
	subStr := s[left:right]
	return subStr
}
