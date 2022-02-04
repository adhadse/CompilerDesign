package parser

import (
	"testing"
)

func TestFirstOfNonTerminal(t *testing.T) {
	testProductionString := []string{
		"E=TD",
		"D=+TD",
		"D=ε",
		"T=FS",
		"S=*FS",
		"S=ε",
		"F=(E)",
		"F=a",
	}
	var testProductionSet = make([][]rune, len(testProductionString))
	for i := range testProductionString {
		testProductionSet[i] = []rune(testProductionString[i])
	}
	if result := FirstOfNonTerminal('E', testProductionSet); !compareRuneSlice(result, []rune{'(', 'a'}) {
		t.Errorf("Method returned other than expected slice for 'E' { (, a }: %v", toStringSlice(result))
	}
	if result := FirstOfNonTerminal('T', testProductionSet); !compareRuneSlice(result, []rune{'(', 'a'}) {
		t.Errorf("Method returned other than expected slice for 'T' { (, a }: %v", toStringSlice(result))
	}
	if result := FirstOfNonTerminal('F', testProductionSet); !compareRuneSlice(result, []rune{'(', 'a'}) {
		t.Errorf("Method returned other than expected slice for 'F'{ (, a }: %v", toStringSlice(result))
	}
	if result := FirstOfNonTerminal('D', testProductionSet); !compareRuneSlice(result, []rune{'+', 'ε'}) {
		t.Errorf("Method returned other than expected slice for D { +, ε }: %v", toStringSlice(result))
	}
	if result := FirstOfNonTerminal('S', testProductionSet); !compareRuneSlice(result, []rune{'*', 'ε'}) {
		t.Errorf("Method returned other than expected slice for S { *, ε }: %v", toStringSlice(result))
	}

}
