package main

import (
	"testing"
)

func TestAbsInt(t *testing.T) {
	if 7 != AbsInt(-7) {
		t.Errorf("failed -7")
	}
	if 105 != AbsInt(105) {
		t.Errorf("failed 105")
	}
}

func TestAtom(t *testing.T) {
	if 7 != Atom(" 7 ") {
		t.Errorf("failed to parse int")
	}
	if " s " != Atom(" s ") {
		t.Errorf("failed to parse string")
	}
}

func TestAtoms(t *testing.T) {
	result := Atoms(" 7,foo   bar: 9.8 ")
	i      := result[0].(int)
	s1     := result[1].(string)
	s2     := result[2].(string)
	f      := result[3].(float64)

	if i != 7 {
		t.Errorf("failed to parse 7")
	}

	if s1 != "foo" {
		t.Errorf("failed to parse foo")
	}

	if s2 != "bar" {
		t.Errorf("failed to parse bar")
	}

	if f < 9.799999 || f > 9.800001 {
		t.Errorf("failed to parse 9.8")
	}
}

func TestDigits(t *testing.T) {
	result := Digits("3.14 is pi, 2.718 is e")
	answer := []int{ 3, 1, 4, 2, 7, 1, 8 }

	for i, val := range answer {
		if result[i] != val {
			t.Errorf("failed to parse digits")
		}
	}
}

func TestInts(t *testing.T) {
	result := Ints("3.14 is pi, 2.718 is e")
	answer := []int{ 3, 14, 2, 718 }

	for i, val := range answer {
		if result[i] != val {
			t.Errorf("failed to parse ints")
		}
	}
}

func TestParse(t *testing.T) {
	result := Parse[int](1, Ints, "\n")

	if result[0][0] != 58789 {
		t.Errorf("foo")
	}
}
