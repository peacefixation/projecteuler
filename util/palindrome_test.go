package util

import "testing"

func TestIsPalindrome(t *testing.T) {
	s := "moom"
	if !IsPalindrome(s) {
		t.Error("Expected true, have false")
	}
}

func TestIsPalindrome2(t *testing.T) {
	s := "906609"
	if !IsPalindrome(s) {
		t.Error("Expected true, have false")
	}
}

func TestIsPalindrome3(t *testing.T) {
	s := "able was I ere I saw elba"
	if !IsPalindrome(s) {
		t.Error("Expected true, have false")
	}
}

func TestIsPalindrome4(t *testing.T) {
	s := "amanaplanacanalpanama"
	if !IsPalindrome(s) {
		t.Error("Expected true, have false")
	}
}
