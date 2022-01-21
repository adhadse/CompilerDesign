package CompilerDesign

func IsValidKeyword(s string) bool {
	if s == TokContinue ||
		s == TokIf ||
		s == TokElse ||
		s == TokBreak ||
		s == TokFor {
		return true
	}
	return false
}
