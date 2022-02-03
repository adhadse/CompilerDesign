package lexical_validators

import "fmt"

func Analyze(s string) {
	left, right := 0, 0
	str := []rune(s)
	length := len(str)

	for right < length && left <= right {
		if IsDelimiter(str[right]) == false {
			right++
		}
		if IsDelimiter(str[right]) && left == right {
			if IsOperator(str[right]) {
				fmt.Printf("'%c' IS AN OPERATOR\n", str[right])
			}
			right++
			left = right
		} else if IsDelimiter(str[right]) && left != right ||
			right == length && left != right {
			subStr := SubString(s, left, right)

			if IsKeyword(subStr) {
				fmt.Printf("'%s' IS A KEYWORD\n", subStr)
			} else if IsInteger(subStr) {
				fmt.Printf("'%s' IS AN INTEGER\n", subStr)
			} else if IsRealNumber(subStr) {
				fmt.Printf("'%s' IS A REAL NUMBER\n", subStr)
			} else if IsIdentifier(subStr) {
				fmt.Printf("'%s' IS A VALID IDENTIFIER\n", subStr)
			} else if IsIdentifier(subStr) == false {
				fmt.Printf("'%s' IS NOT A VALID IDENTIFIER\n", subStr)
			}
			left = right
		}
	}
	return
}
