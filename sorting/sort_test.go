package sorting

import "testing"

func TestSortAscending(t *testing.T) {
	l1 := []int{10, 2, 15, 24, 32}
	e1 := []int{2, 10, 15, 24, 32}

	SortAscending(l1)

	if !listsequal(l1, e1) {
		t.Errorf("SortAscending(%v) returned %d, but expected %d", l1, l1, e1)
	}
}

func TestSortStringsByLengthAscending(t *testing.T) {
	l1 := []string{"John", "Mathilda", "Bob"}
	e1 := []string{"Bob", "John", "Mathilda"}

	SortStringsByLengthAscending(l1)

	if !listsequal(l1, e1) {
		t.Errorf("SortStringsByLengthAscending(%v) returned %v, but expected %v", l1, l1, e1)
	}

}

func TestSortByAgeDescending(t *testing.T) {
	l1 := []Person{
		{"John", 20},
		{"Jane", 10},
		{"Bob", 30},
	}
	e1 := []Person{
		{"Bob", 30},
		{"John", 20},
		{"Jane", 10},
	}

	SortByAgeDescending(l1)

	if !listsequal(l1, e1) {
		t.Errorf("SortByAgeDescending(%v) returned %v, but expected %v", l1, l1, e1)
	}
}

func TestSortByNameAscending(t *testing.T) {
	l1 := []Person{
		{"John", 20},
		{"Jane", 10},
		{"Bob", 30},
	}
	e1 := []Person{
		{"Bob", 30},
		{"Jane", 10},
		{"John", 20},
	}

	SortByNameAscending(l1)

	if !listsequal(l1, e1) {
		t.Errorf("SortByAgeAscending(%v) returned %v, but expected %v", l1, l1, e1)
	}
}
