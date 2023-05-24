package mapfunction

import "testing"

func TestMapToString(t *testing.T) {
	l1 := []int{10, 2, 15, 24, 32}
	e1 := []string{"10", "2", "15", "24", "32"}

	r1 := MapToString(l1)

	if !listsequal(r1, e1) {
		t.Errorf("MapToString(%v) returned %v, but expected %v", l1, r1, e1)
	}
}

func TestDoubleEach(t *testing.T) {
	l1 := []int{10, 2, 15, 24, 32}
	e1 := []int{20, 4, 30, 48, 64}

	r1 := DoubleEach(l1)

	if !listsequal(r1, e1) {
		t.Errorf("DoubleEach(%v) returned %v, but expected %v", l1, r1, e1)
	}
}

func TestAppendToEach(t *testing.T) {
	l1 := []string{"John", "Jane", "Bob"}
	e1 := []string{"John Doe", "Jane Doe", "Bob Doe"}

	r1 := AppendToEach(l1, " Doe")

	if !listsequal(r1, e1) {
		t.Errorf("AppendToEach(%v) returned %v, but expected %v", l1, r1, e1)
	}
}
