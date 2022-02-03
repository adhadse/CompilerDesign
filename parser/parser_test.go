package parser

import (
	"fmt"
	"testing"
)

func TestFirstOfNonTerminal(t *testing.T) {
	var result []rune
	testProductionString := []string{
		"E=TD",
		"D=+TD",
		"D=$",
		"T=FS",
		"S=*FS",
		"S=$",
		"F=E",
		"F=a",
	}
	var testProductionSet = make([][]rune, len(testProductionString))
	for i := range testProductionString {
		testProductionSet[i] = []rune(testProductionString[i])
	}
	result = FirstOfNonTerminal(result, 'S', testProductionSet)
	fmt.Println(string(result))
}
