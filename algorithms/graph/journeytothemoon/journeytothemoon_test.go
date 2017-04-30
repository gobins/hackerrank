package main

import "testing"

func TestInitAdjMat(t *testing.T) {
	var q = []pair{
		{
			a: 0,
			b: 1,
		},
		{
			a: 2,
			b: 3,
		},
		{
			a: 0,
			b: 4,
		},
	}

	output := initAdjMat(q)

	c1 := output[0]
	if !(contains(1, c1.neighbours)) {
		t.Error("Adjacent matrix is not correct")
	}
	if !(contains(4, c1.neighbours)) {
		t.Error("Adjacent matrix is not correct")
	}

	c2 := output[1]
	if !(contains(0, c2.neighbours)) {
		t.Error("Adjacent matrix is not correct")
	}
	if contains(4, c2.neighbours) {
		t.Error("Adjacent matrix is not correct")
	}

	c3 := output[3]
	if !(contains(2, c3.neighbours)) {
		t.Error("Adjacent matrix is not correct")
	}
	if contains(0, c3.neighbours) {
		t.Error("Adjacent matrix is not correct")
	}

}

func contains(v int, n []int) bool {
	var flag bool
	for _, x := range n {
		if x == v {
			return true
		}
	}
	return flag
}
