package parser

import "unicode"

// FirstOfNonTerminal finds the FIRST set; the set of terminals you could
// possibly see as the first part of the expansion of a non-terminal used in predictive parsers
//
// The rules for finding FIRST of a given grammar is :
//
// 1. If X is terminal, then FIRST(X) = {X}.
// 2. If  X-> ε is a production, then First(X) = First(X) + ε
// 3. If X is a non-terminal and X -> Y_1 Y_2...Y_k is a production,
//    then if Y_1 doesn’t derive epsilon all in Y_1 will come under X else move to Y_2 and so on.
func FirstOfNonTerminal(result []rune, c rune, productionSet [][]rune) []rune {
	var subResult []rune
	numOfProductions := len(productionSet)
	// If X is terminal, First(X) = {X}
	if unicode.IsLower(c) {
		result = addToResultSet(result, c)
		return result
	}
	// If X is non-terminal,
	// then read each production
	for i := 0; i < numOfProductions; i++ {
		// Find production with X as LHS
		if productionSet[i][0] == c {
			// If X -> ε is a production, then add ε to FIRST(X)
			if productionSet[i][2] == '$' {
				result = addToResultSet(result, '$')
			} else {
				// If X is a non-terminal, and X -> Y1 Y2 ... Yk
				// is a production, then add 'a' to FIRST(X)
				// if for some i, a is in FIRST (Yi),
				// and ε is in all of FIRST(Y1), ..., FIRST(Yi-1)
				j := 2
				for ; j < len(productionSet[i]); j++ {
					foundEpsilon := false
					subResult = FirstOfNonTerminal(subResult, productionSet[i][j], productionSet)
					for k := 0; k < len(subResult); k++ {
						result = addToResultSet(result, subResult[k])
					}
					for k := 0; k < len(subResult); k++ {
						if subResult[k] == '$' {
							foundEpsilon = true
							break
						}
					}
					// No ε found, no need to check next element
					if !foundEpsilon {
						break
					}
				}
			}
		}
	}
	return result
}

// addToResultSet adds the computed element to result set
// This function is useful instead of built-in append as it
// avoids multiple inclusion of elements
func addToResultSet(result []rune, val rune) []rune {
	for i := 0; i < len(result); i++ {
		if result[i] == val {
			return result
		}
	}
	result = append(result, val)
	return result
}
