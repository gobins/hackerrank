package main

import "testing"

func TestToInt(t *testing.T) {
	a := "24"
	b := "99"
	if toInt(a) != 24 {
		t.Error("Error converting string to int")
	}

	if toInt(b) != 99 {
		t.Error("Error converting string to int")
	}
}
