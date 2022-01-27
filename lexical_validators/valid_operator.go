package lexical_validators

// IsOperator returns 'true' if the character/rune is an OPERATOR.
func IsOperator(ch rune) bool {
	if ch == TokAddition || ch == TokSubtraction || ch == TokMultiplication ||
		ch == TokDivision || ch == TokGreaterThan || ch == TokLessThan ||
		ch == TokEquals {
		return true
	}
	return false
}
