package CompilerDesign

func IsValidIdentifier(s string) bool {
	// Check if string containing single word, without any spaces
	// is identifier or not by checking if between a and Z or A and Z
	// or it starts with underscore.
	// Also make sure it is not a valid keyword
	str := []rune(s)
	if IsValidKeyword(s) {
		return false
	}
	// if first character is invalid
	if !(str[0] >= TOK_SMALL_A && str[0] <= TOK_SMALL_Z ||
		str[0] >= TOK_CAP_A && str[0] <= TOK_CAP_Z ||
		str[0] == TOK_UNDERSCORE) {
		return false
	}

	// Traverse the string for rest of characters
	for i := 1; i < len(str); i++ {
		if !(str[i] >= TOK_SMALL_A && str[i] <= TOK_SMALL_Z ||
			str[i] >= TOK_CAP_A && str[i] <= TOK_CAP_Z ||
			str[i] >= TOK_0 && str[i] <= TOK_9 ||
			str[i] == TOK_UNDERSCORE) {
			return false
		}
	}
	return true
}

func IsValidKeyword(s string) bool {
	if s == TOK_CONTINUE ||
		s == TOK_IF ||
		s == TOK_ELSE ||
		s == TOK_BREAK ||
		s == TOK_FOR {
		return true
	}
	return false
}
