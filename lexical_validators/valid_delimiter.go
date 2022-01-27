package lexical_validators

// IsDelimiter returns 'true' if the character/rune is a DELIMITER.
func IsDelimiter(ch rune) bool {
	if ch == TokSpace || ch == TokAddition || ch == TokSubtraction || ch == TokMultiplication ||
		ch == TokDivision || ch == TokComma || ch == TokSemicolon ||
		ch == TokGreaterThan || ch == TokLessThan || ch == TokEquals ||
		ch == TokLeftBracket || ch == TokRightBracket || ch == TokSqLeftBracket ||
		ch == TokSqRightBracket || ch == TokCurLeftBracket || ch == TokCurRightBracket {
		return true
	}
	return false
}
