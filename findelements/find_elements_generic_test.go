package findelements

import "testing"

func TestListContains_int(t *testing.T) {
	l1 := []int{1, 2, 3, 4, 5}
	r1_2 := ListContains(l1, 2)
	r1_17 := ListContains(l1, 17)

	if !r1_2 {
		t.Errorf("ListContains(%v, %d) returned %t, but expected %t", l1, 2, r1_2, true)
	}
	if r1_17 {
		t.Errorf("ListContains(%v, %d) returned %t, but expected %t", l1, 17, r1_17, false)
	}
}
func TestListContains_string(t *testing.T) {
	l1 := []string{"a", "b", "c", "d", "e"}
	r1_b := ListContains(l1, "a")
	r1_f := ListContains(l1, "f")

	if !r1_b {
		t.Errorf("ListContains(%v, %d) returned %t, but expected %t", l1, 2, r1_b, true)
	}
	if r1_f {
		t.Errorf("ListContains(%v, %d) returned %t, but expected %t", l1, 17, r1_f, false)
	}
}

func TestListContains_struct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	l1 := []Person{
		{"Alice", 23},
		{"Bob", 42},
		{"Charlie", 17},
	}
	r1_bob := ListContains(l1, Person{"Bob", 42})
	r1_frank := ListContains(l1, Person{"Frank", 17})

	if !r1_bob {
		t.Errorf("ListContains(%v, %v) returned %t, but expected %t", l1, Person{"Bob", 42}, r1_bob, true)
	}
	if r1_frank {
		t.Errorf("ListContains(%v, %v) returned %t, but expected %t", l1, Person{"Frank", 17}, r1_frank, false)
	}
}
