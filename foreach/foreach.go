package foreach

// ForEach führt eine Funktion für jedes Element einer Liste aus.
// Die Funktion f erwartet einen Wert vom Typ T und liefert keinen Wert zurück.
func ForEach[T any](list []T, f func(T)) {
	for _, v := range list {
		f(v)
	}
}

// CountEven erwartet eine Liste von Zahlen.
// CountEven zählt, wie viele Elemente der Liste gerade sind.
func CountEven(list []int) int {
	count := 0
	ForEach(list, func(v int) {
		if v%2 == 0 {
			count++
		}
	})
	return count
}

// CountGreater erwartet eine Liste von Zahlen und eine Zahl x.
// CountGreater zählt, wie viele Elemente der Liste größer als x sind.
func CountGreater(list []int, x int) int {
	count := 0
	ForEach(list, func(v int) {
		if v > x {
			count++
		}
	})
	return count
}
