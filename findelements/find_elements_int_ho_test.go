package findelements

import "testing"

func TestFindFirstNegative(t *testing.T) {
	l1 := []int{1, -2, 3, 4, 5}
	l2 := []int{1, 2, 3, -4, 5}
	l3 := []int{1, 2, 3, 4, 5}

	r1 := FindFirstNegative(l1)
	r2 := FindFirstNegative(l2)
	r3 := FindFirstNegative(l3)

	if r1 != 1 {
		t.Errorf("FindFirstNegative(%v) returned %d, but expected %d", l1, r1, 1)
	}
	if r2 != 3 {
		t.Errorf("FindFirstNegative(%v) returned %d, but expected %d", l2, r2, 3)
	}
	if r3 != -1 {
		t.Errorf("FindFirstNegative(%v) returned %d, but expected %d", l3, r3, -1)
	}
}

func TestFindFirstMultipleOf3(t *testing.T) {
	l1 := []int{1, 15, 2, 4, 8}
	l2 := []int{1, 2, 4, 15, 8}
	l3 := []int{1, 2, 4, 8, 16}

	r1 := FindFirstMultipleOf3(l1)
	r2 := FindFirstMultipleOf3(l2)
	r3 := FindFirstMultipleOf3(l3)

	if r1 != 1 {
		t.Errorf("FindFirstMultipleOf3(%v) returned %d, but expected %d", l1, r1, 1)
	}
	if r2 != 3 {
		t.Errorf("FindFirstMultipleOf3(%v) returned %d, but expected %d", l2, r2, 3)
	}
	if r3 != -1 {
		t.Errorf("FindFirstMultipleOf3(%v) returned %d, but expected %d", l3, r3, -1)
	}
}

func TestFindFirstDivisorOf60(t *testing.T) {
	l1 := []int{7, 15, 21, 28, 35}
	l2 := []int{7, 14, 21, 15, 35}
	l3 := []int{7, 14, 21, 28, 35}

	r1 := FindFirstDivisorOf60(l1)
	r2 := FindFirstDivisorOf60(l2)
	r3 := FindFirstDivisorOf60(l3)

	if r1 != 1 {
		t.Errorf("FindFirstDivisorOf60(%v) returned %d, but expected %d", l1, r1, 1)
	}
	if r2 != 3 {
		t.Errorf("FindFirstDivisorOf60(%v) returned %d, but expected %d", l2, r2, 3)
	}
	if r3 != -1 {
		t.Errorf("FindFirstDivisorOf60(%v) returned %d, but expected %d", l3, r3, -1)
	}
}

func TestFindFirstCommonMultiple(t *testing.T) {
	l1 := []int{7, 15, 21, 28, 35}
	l2 := []int{7, 14, 21, 15, 35}
	l3 := []int{7, 14, 21, 28, 35}

	r1 := FindFirstCommonMultiple(l1, 3, 5)
	r2 := FindFirstCommonMultiple(l2, 3, 5)
	r3 := FindFirstCommonMultiple(l3, 3, 5)

	if r1 != 1 {
		t.Errorf("FindFirstCommonMultiple(%v, 3, 5) returned %d, but expected %d", l1, r1, 1)
	}
	if r2 != 3 {
		t.Errorf("FindFirstCommonMultiple(%v, 3, 5) returned %d, but expected %d", l2, r2, 3)
	}
	if r3 != -1 {
		t.Errorf("FindFirstCommonMultiple(%v, 3, 5) returned %d, but expected %d", l3, r3, -1)
	}
}
