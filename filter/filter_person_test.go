package filter

import "testing"

// TestFilterMinAge prüft, ob FilterMinAge korrekt funktioniert.
func TestFilterMinAge(t *testing.T) {
	list := []Person{
		{"Alice", 23},
		{"Bob", 17},
		{"Charlie", 42},
		{"Diana", 13},
	}

	r := FilterMinAge(list, 18)
	e := []Person{
		{"Alice", 23},
		{"Charlie", 42},
	}

	if !listsequal(r, e) {
		t.Errorf("FilterMinAge(%v) returned %v, but expected %v", list, r, e)
	}
}

// TestFilterLongNames prüft, ob FilterLongNames korrekt funktioniert.
func TestFilterLongNames(t *testing.T) {
	list := []Person{
		{"Alice", 23},
		{"Bob", 17},
		{"Charlie", 42},
		{"Diana", 13},
	}

	r := FilterLongNames(list, 5)
	e := []Person{
		{"Alice", 23},
		{"Charlie", 42},
		{"Diana", 13},
	}

	if !listsequal(r, e) {
		t.Errorf("FilterLongNames(%v) returned %v, but expected %v", list, r, e)
	}
}

// TestFilterNamePrefix prüft, ob FilterNamePrefix korrekt funktioniert.
func TestFilterNamePrefix(t *testing.T) {
	list := []Person{
		{"Alice", 23},
		{"Bob", 17},
		{"Charlie", 42},
		{"Diana", 13},
	}

	r := FilterNamePrefix(list, "C")
	e := []Person{
		{"Charlie", 42},
	}

	if !listsequal(r, e) {
		t.Errorf("FilterNamePrefix(%v) returned %v, but expected %v", list, r, e)
	}
}

// TestFilterChildren prüft, ob FilterChildren korrekt funktioniert.
func TestFilterChildren(t *testing.T) {
	list := []Person{
		{"Alice", 23},
		{"Bob", 17},
		{"Charlie", 42},
		{"Diana", 13},
	}

	r := FilterChildren(list)
	e := []Person{
		{"Diana", 13},
	}

	if !listsequal(r, e) {
		t.Errorf("FilterChildren(%v) returned %v, but expected %v", list, r, e)
	}
}

// TestFilterChildrenWithLongNames prüft, ob FilterChildrenWithLongNames korrekt funktioniert.
func TestFilterChildrenWithLongNames(t *testing.T) {
	list := []Person{
		{"Alice", 23},
		{"Bob", 17},
		{"Charlie", 42},
		{"Diana", 13},
		{"Emma", 12},
	}

	r := FilterChildrenWithLongNames(list, 5)
	e := []Person{
		{"Diana", 13},
	}

	if !listsequal(r, e) {
		t.Errorf("FilterChildrenWithLongNames(%v) returned %v, but expected %v", list, r, e)
	}
}
