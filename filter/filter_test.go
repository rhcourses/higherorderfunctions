package filter

import "testing"

func TestFilterSmaller(t *testing.T) {
	l1 := []int{10, 2, 15, 24, 32}

	r1 := FilterSmaller(l1, 20)
	e1 := []int{10, 2, 15}

	if !listsequal(r1, e1) {
		t.Errorf("FilterSmaller(%v) returned %d, but expected %d", l1, r1, e1)
	}
}
