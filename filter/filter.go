package filter

// FilterList erwartet eine Liste und ein Prädikat
// (eine Funktion, die einen Wert nimmt und true oder false zurückgibt).
// FilterList liefert eine neue Liste, die nur die Elemente der ursprünglichen Liste enthält,
// für die das Prädikat true liefert.
func FilterList[T any](list []T, predicate func(T) bool) []T {
	result := []T{}
	for _, v := range list {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// FilterSmaller erwartet eine Liste und eine Zahl x.
// FilterSmaller liefert eine neue Liste, die nur die Elemente der ursprünglichen Liste enthält,
// die kleiner als x sind.
func FilterSmaller(list []int, x int) []int {
	return FilterList(list, func(v int) bool {
		return v < x
	})
}
