package CompilerDesign

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
