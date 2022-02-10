package lexical_analyzer

import (
	"testing"
)

func TestIsIdentifier(t *testing.T) {
	if val := IsIdentifier("12"); val != false {
		t.Error("Method returned true for identifier: 12")
	}
	if val := IsIdentifier("bear"); val != true {
		t.Error("Method returned false for identifier: bear")
	}
	if val := IsIdentifier("break"); val != false {
		t.Error("Method returned true for identifier: bear")
	}
}

func TestIsKeyword(t *testing.T) {
	if val := IsKeyword("12"); val != false {
		t.Error("Method returned true for keyword: 12")
	}
	if val := IsKeyword("break"); val != true {
		t.Error("Method returned false for keyword: break")
	}
}

func TestIsOperator(t *testing.T) {
	if val := IsOperator('/'); val != true {
		t.Error("Method returned true for : /")
	}
	if val := IsOperator('+'); val != true {
		t.Error("Method returned true for operator: +")
	}
	if val := IsOperator('b'); val != false {
		t.Error("Method returned false for identifier: b")
	}
}

func TestIsDelimiter(t *testing.T) {
	if val := IsDelimiter(' '); val != true {
		t.Error("Method returned false for empty space")
	}
	if val := IsDelimiter('('); val != true {
		t.Error("Method returned false for delimiter: (")
	}
	if val := IsDelimiter('b'); val != false {
		t.Error("Method returned true for identifier: ,")
	}
}

func TestIsInteger(t *testing.T) {
	if val := IsInteger("1"); val != true {
		t.Error("Method returned false for integer: 1")
	}
	if val := IsInteger("123"); val != true {
		t.Error("Method returned false for integer: 123")
	}
	if val := IsInteger(""); val != false {
		t.Error("Method returned true for empty space")
	}
	if val := IsInteger("123-4"); val != false {
		t.Error("Method returned true for: 123-4")
	}
}

func TestIsRealNumber(t *testing.T) {
	if val := IsRealNumber("1"); val != false {
		t.Error("Method returned true for integer: 1")
	}
	if val := IsRealNumber("123"); val != false {
		t.Error("Method returned true for integer: 123")
	}
	if val := IsRealNumber("123."); val != true {
		t.Error("Method returned false for realnumber: 123.")
	}
	if val := IsRealNumber(""); val != false {
		t.Error("Method returned true for empty space")
	}
	if val := IsRealNumber("123-4"); val != false {
		t.Error("Method returned true for: 123-4")
	}
}

func TestSubString(t *testing.T) {
	if val := SubString("a happy", 0, 1); val != "a" {
		t.Error("Method returned other than expected value for: \"a happy\"", val)
	}
	if val := SubString("university", 4, 10); val != "ersity" {
		t.Error("Method returned other than expected value for: milkcocoa", val)
	}
}

func TestAnalyze(t *testing.T) {
	Analyze("for i=1; i<10;i++;")
}
