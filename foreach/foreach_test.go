package foreach

import "testing"

func TestCountEven(t *testing.T) {
	l1 := []int{1, 2, 3, 4, 5}
	l2 := []int{1, 3, 5, 7, 2, 9}
	l3 := []int{1, 3, 5, 7, 9}

	r1 := CountEven(l1)
	r2 := CountEven(l2)
	r3 := CountEven(l3)

	if r1 != 2 {
		t.Errorf("CountEven(%v) returned %d, but expected %d", l1, r1, 2)
	}
	if r2 != 1 {
		t.Errorf("CountEven(%v) returned %d, but expected %d", l2, r2, 1)
	}
	if r3 != 0 {
		t.Errorf("CountEven(%v) returned %d, but expected %d", l3, r3, 0)
	}
}

func TestCountGreater(t *testing.T) {
	l1 := []int{1, 2, 3, 4, 5}
	l2 := []int{1, 2, 3, 4, 5, 6}
	l3 := []int{1, 2, 3}

	r1 := CountGreater(l1, 3)
	r2 := CountGreater(l2, 3)
	r3 := CountGreater(l3, 3)

	if r1 != 2 {
		t.Errorf("CountGreater(%v) returned %d, but expected %d", l1, r1, 2)
	}
	if r2 != 3 {
		t.Errorf("CountGreater(%v) returned %d, but expected %d", l2, r2, 3)
	}
	if r3 != 0 {
		t.Errorf("CountGreater(%v) returned %d, but expected %d", l3, r3, 0)
	}
}
