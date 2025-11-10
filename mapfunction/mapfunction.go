package mapfunction

import "fmt"

// MapFunction erwartet eine Liste von Werten und eine Funktion.
// MapFunction wendet die Funktion auf alle Elemente der Liste an und
// liefert eine neue Liste mit den Ergebnissen.
func MapFunction[T, U any](list []T, f func(T) U) []U {
	result := []U{}
	for _, v := range list {
		result = append(result, f(v))
	}
	return result
}

// MapToString erwartet eine Liste von Werten.
// MapToString wendet die Funktion fmt.Sprint auf alle Elemente der Liste an und
// liefert eine neue Liste mit den Ergebnissen.
func MapToString[T any](list []T) []string {
	return MapFunction(list, func(v T) string {
		return fmt.Sprint(v)
	})
}

// DoubleEach erwartet eine Liste von Zahlen.
// DoubleEach verdoppelt alle Elemente der Liste und liefert eine neue Liste mit den Ergebnissen.
func DoubleEach(list []int) []int {
	// HINWEIS:
	// Rufen Sie die Funktion MapFunction auf und übergeben Sie ihr eine Funktion,
	// die ein Element vom Typ int erwartet und dieses Element verdoppelt.
	// solution:begin
	return MapFunction(list, func(v int) int {
		return v * 2
	})
	// solution:end
}

// AppendToEach erwartet eine Liste von Strings und einen String s.
// AppendToEach hängt s an alle Elemente der Liste an und liefert eine neue Liste mit den Ergebnissen.
func AppendToEach(list []string, s string) []string {
	// HINWEIS:
	// Rufen Sie die Funktion MapFunction auf und übergeben Sie ihr eine Funktion,
	// die ein Element vom Typ string erwartet und an dieses Element den String s anhängt.
	// solution:begin
	return MapFunction(list, func(v string) string {
		return v + s
	})
	// solution:end
}
