package findelements

import "testing"

func TestFindFirstEven(t *testing.T) {
	l1 := []int{1, 2, 3, 4, 5}
	l2 := []int{1, 3, 5, 7, 2, 9}
	l3 := []int{1, 3, 5, 7, 9}

	r1 := FindFirstEven(l1)
	r2 := FindFirstEven(l2)
	r3 := FindFirstEven(l3)

	if r1 != 1 {
		t.Errorf("FindFirstEven(%v) returned %d, but expected %d", l1, r1, 1)
	}
	if r2 != 4 {
		t.Errorf("FindFirstEven(%v) returned %d, but expected %d", l2, r2, 4)
	}
	if r3 != -1 {
		t.Errorf("FindFirstEven(%v) returned %d, but expected %d", l3, r3, -1)
	}
}

func TestFindFirstOdd(t *testing.T) {
	l1 := []int{1, 2, 3, 4, 5}
	l2 := []int{2, 4, 6, 8, 1, 10}
	l3 := []int{2, 4, 6, 8, 10}

	r1 := FindFirstOdd(l1)
	r2 := FindFirstOdd(l2)
	r3 := FindFirstOdd(l3)

	if r1 != 0 {
		t.Errorf("FindFirstOdd(%v) returned %d, but expected %d", l1, r1, 0)
	}
	if r2 != 4 {
		t.Errorf("FindFirstOdd(%v) returned %d, but expected %d", l2, r2, 4)
	}
	if r3 != -1 {
		t.Errorf("FindFirstOdd(%v) returned %d, but expected %d", l3, r3, -1)
	}
}

func TestFindFirstGreater3(t *testing.T) {
	l1 := []int{1, 2, 3, 4, 5}
	l2 := []int{1, 2, 3, 4, 5, 6}
	l3 := []int{1, 2, 3}

	r1 := FindFirstGreater3(l1)
	r2 := FindFirstGreater3(l2)
	r3 := FindFirstGreater3(l3)

	if r1 != 3 {
		t.Errorf("FindFirstGreater3(%v) returned %d, but expected %d", l1, r1, 3)
	}
	if r2 != 3 {
		t.Errorf("FindFirstGreater3(%v) returned %d, but expected %d", l2, r2, 3)
	}
	if r3 != -1 {
		t.Errorf("FindFirstGreater3(%v) returned %d, but expected %d", l3, r3, -1)
	}
}
