package sorting

import "sort"

// SortAscending sortiert die Liste aufsteigend.
// Die Funktion verwendet dafür die Funktion sort.Slice aus dem Paket sort.
// Diese erwartet eine Liste und eine Funktion, die zwei Elemente der Liste vergleicht.
// Die Vergleichsfunktion (ein *Komparator*) erwartet zwei Position und liefert true,
// falls das Element an der ersten Position kleiner ist als das Element an der zweiten Position.
func SortAscending(list []int) {
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
}

// SortStringsByLengthAscending erwartet eine Liste von Strings.
// SortStringsByLengthAscending sortiert die Liste nach der Länge der Strings aufsteigend.
// Die Funktion verwendet dafür die Funktion sort.Slice aus dem Paket sort.
func SortStringsByLengthAscending(list []string) {
	sort.Slice(list, func(i, j int) bool {
		return len(list[i]) < len(list[j])
	})
}

// Person ist (mal wieder) ein Datentyp für Personen.
type Person struct {
	Name string
	Age  int
}

// SortByAgeDescending sortiert die Liste nach dem Alter der Personen absteigend.
// Die Funktion verwendet dafür die Funktion sort.Slice aus dem Paket sort.
func SortByAgeDescending(list []Person) {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Age > list[j].Age
	})
}

// SortByNameAscending sortiert die Liste nach dem Namen der Personen aufsteigend.
// Die Funktion verwendet dafür die Funktion sort.Slice aus dem Paket sort.
func SortByNameAscending(list []Person) {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})
}
