package CompilerDesign

import "testing"

func TestIsValidIdentifier(t *testing.T) {
	if val := IsValidIdentifier("12"); val != false {
		t.Error("Method returned true for identifier: 12")
	}
	if val := IsValidIdentifier("bear"); val != true {
		t.Error("Method returned false for identifier: bear")
	}
	if val := IsValidIdentifier("break"); val != false {
		t.Error("Method returned true for identifier: bear")
	}
}

func TestIsValidKeyword(t *testing.T) {
	if val := IsValidKeyword("12"); val != false {
		t.Error("Method returned true for keyword: 12")
	}
	if val := IsValidKeyword("break"); val != true {
		t.Error("Method returned false for keyword: break")
	}
}
