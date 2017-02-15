package main

import "testing"

func TestCheckContains(t *testing.T) {
	a := "abcd"
	flag := checkContains(a)
	if flag {
		t.Error("String does not contain repeated characters")
	}

	b := "aba"
	flag = checkContains(b)
	if !flag {
		t.Error("String contains repeated characters")
	}
}
