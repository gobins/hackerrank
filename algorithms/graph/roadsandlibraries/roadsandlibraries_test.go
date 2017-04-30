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

func TestInitAdjMat(t *testing.T) {
	var q = Query{
		m: 3,
		n: 3,
		connections: []*Connect{
			{
				u: 1,
				v: 2,
			},
			{
				u: 3,
				v: 1,
			},
			{
				u: 2,
				v: 3,
			},
		},
	}

	output := q.initAdjMat()

	c1 := output[1]
	if !(contains(2, c1.neighbours)) {
		t.Error("Adjacent matrix is not correct")
	}
	if !(contains(3, c1.neighbours)) {
		t.Error("Adjacent matrix is not correct")
	}

	c2 := output[2]
	if !(contains(1, c2.neighbours)) {
		t.Error("Adjacent matrix is not correct")
	}
	if !(contains(3, c2.neighbours)) {
		t.Error("Adjacent matrix is not correct")
	}

	c3 := output[3]
	if !(contains(1, c3.neighbours)) {
		t.Error("Adjacent matrix is not correct")
	}
	if !(contains(2, c3.neighbours)) {
		t.Error("Adjacent matrix is not correct")
	}

}

func TestProcessQuery(t *testing.T) {
	var q1 = Query{
		m:     3,
		n:     3,
		clib:  2,
		croad: 1,
		connections: []*Connect{
			{
				u: 1,
				v: 2,
			},
			{
				u: 3,
				v: 1,
			},
			{
				u: 2,
				v: 3,
			},
		},
	}

	c := q1.processQuery()

	if c != 4 {
		t.Errorf("Calculated Cost %d is not equal to expected cost %d", c, 4)
	}

	var q2 = Query{
		m:     3,
		n:     3,
		clib:  2,
		croad: 3,
		connections: []*Connect{
			{
				u: 1,
				v: 2,
			},
			{
				u: 3,
				v: 1,
			},
			{
				u: 2,
				v: 3,
			},
		},
	}

	c = q2.processQuery()

	if c != 6 {
		t.Errorf("Calculated Cost %d is not equal to expected cost %d", c, 6)
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
