package main

import "testing"

func TestIsBlankStringEmptyString(t *testing.T) {
	result := isBlankString("")
	if result == false {
		t.Errorf("isBlankString('') = %t; want true", result)
	}
}

func TestIsBlankStringWhitespaceString(t *testing.T) {
	result := isBlankString("   ")
	if result == false {
		t.Errorf("isBlankString('   ' = %t; want true", result)
	}
}

func TestIsBlankStringNonEmptyString(t *testing.T) {
	result := isBlankString("k")
	if result == true {
		t.Errorf("isBlankString('k') = %t; want false", result)
	}
}
