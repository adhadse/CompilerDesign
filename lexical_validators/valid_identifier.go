package lexical_validators

// IsIdentifier check if string containing single word, without any spaces
// is identifier or not by checking if between a and Z or A and Z
// or it starts with underscore.
// Also make sure it is not a valid keyword
func IsIdentifier(s string) bool {
	str := []rune(s)
	if IsKeyword(s) {
		return false
	}
	// if first character is invalid
	if !(str[0] >= TokSmallA && str[0] <= TokSmallZ ||
		str[0] >= TokCapA && str[0] <= TokCapZ ||
		str[0] == TokUnderscore) {
		return false
	}

	// Traverse the string for rest of characters
	for i := 1; i < len(str); i++ {
		if !(str[i] >= TokSmallA && str[i] <= TokSmallZ ||
			str[i] >= TokCapA && str[i] <= TokCapZ ||
			str[i] >= Tok0 && str[i] <= Tok9 ||
			str[i] == TokUnderscore) {
			return false
		}
	}
	return true
}
