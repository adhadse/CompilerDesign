package lexical_validators

// IsRealNumber returns 'true' if the string is a REAL NUMBER.
func IsRealNumber(s string) bool {
	str := []rune(s)
	length := len(str)
	hasDecimal := false

	if length == 0 {
		return false
	}
	for i := 0; i < length; i++ {
		if str[i] != Tok0 && str[i] != Tok1 && str[i] != Tok2 &&
			str[i] != Tok3 && str[i] != Tok4 && str[i] != Tok5 &&
			str[i] != Tok6 && str[i] != Tok7 && str[i] != Tok8 &&
			str[i] != Tok9 && str[i] != TokDecimal ||
			(str[i] == TokDash && i > 0) {
			return false
		}
		if str[i] == TokDecimal {
			hasDecimal = true
		}
	}
	return hasDecimal
}
