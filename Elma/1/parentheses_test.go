package main

import (
	"testing"
)

func TestParenthesesTask(t *testing.T) {
	case1 := "{}"
	expectedFor1 := "YES"

	result := ParenthesesTask(case1)

	if result != expectedFor1 {
		t.Errorf("Wrong result. Expected: %s, got: %s ", expectedFor1, result)
	}

	case2 := "{312312}[[]1231)"
	expectedFor2 := "НЕТ"

	result = ParenthesesTask(case2)

	if result != expectedFor2 {
		t.Errorf("Wrong result. Expected: %s, got: %s ", expectedFor2, result)
	}

	case3 := "(){}[]"
	expectedFor3 := "YES"

	result = ParenthesesTask(case3)

	if result != expectedFor3 {
		t.Errorf("Wrong result. Expected: %s, got: %s ", expectedFor3, result)
	}

	case4 := "([])"
	expectedFor4 := "YES"

	result = ParenthesesTask(case4)

	if result != expectedFor4 {
		t.Errorf("Wrong result. Expected: %s, got: %s ", expectedFor4, result)
	}

	case5 := ""
	expectedFor5 := "YES"

	result = ParenthesesTask(case5)

	if result != expectedFor5 {
		t.Errorf("Wrong result. Expected: %s, got: %s ", expectedFor5, result)
	}

	case6 := "123321312"
	expectedFor6 := "YES"

	result = ParenthesesTask(case6)

	if result != expectedFor6 {
		t.Errorf("Wrong result. Expected: %s, got: %s ", expectedFor6, result)
	}
}
