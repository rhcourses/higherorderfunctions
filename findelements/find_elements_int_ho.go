package findelements

// FindFirst liefert die Position der ersten Zahl in der Liste, die die Funktion f erfüllt.
// Falls keine solche Zahl gefunden wird, wird -1 zurückgegeben.
func FindFirst(list []int, f func(int) bool) int {
	for i, v := range list {
		if f(v) {
			return i
		}
	}
	return -1
}

// FindFirstNegative liefert die Position der ersten negativen Zahl in der Liste.
// Falls keine negative Zahl gefunden wird, wird -1 zurückgegeben.
//
// Anmerkung: Die Funktion FindFirstNegative benuzt die Funktion FindFirst.
// D.h. sie ruft die Funktion FindFirst auf und übergibt ihr eine Funktion als Parameter.
// Diese Funktion ist weiter unten mit dem Namen "negative" definiert.
func FindFirstNegative(list []int) int {
	return FindFirst(list, negative)
}

// negative erwartet eine Zahl und prüft, ob diese negativ ist.
func negative(x int) bool {
	return x < 0
}

// FindFirstMultipleOf3 liefert die Position der ersten Zahl in der Liste, die ein Vielfaches von 3 ist.
// Falls keine solche Zahl gefunden wird, wird -1 zurückgegeben.
//
// Anmerkung: Die Funktion FindFirstMultipleOf3 benuzt die Funktion FindFirst.
// D.h. sie ruft die Funktion FindFirst auf und übergibt ihr eine Funktion als Parameter.
// Hier gehen wir noch einen Schritt weiter als bei FindFirstNegative:
// Statt die Funktion weiter unten separat zu definieren, definieren wir sie direkt hier
// innerhalb der Funktion FindFirstMultipleOf3.
func FindFirstMultipleOf3(list []int) int {
	ismultipleof3 := func(x int) bool {
		return x%3 == 0
	}
	return FindFirst(list, ismultipleof3)
}

// FindFirstDivisorOf60 liefert die Position der ersten Zahl in der Liste, die ein Teiler von 60 ist.
// Falls keine solche Zahl gefunden wird, wird -1 zurückgegeben.
//
// Anmerkung: Wie bei FindFirstMultipleOf3 definieren wir die Funktion direkt hier.
// Wir geben ihr aber nicht einmal mehr einen Namen, sondern definieren eine *anonyme Funktion*.
func FindFirstDivisorOf60(list []int) int {
	return FindFirst(list, func(x int) bool {
		return 60%x == 0
	})
}

// FindFirstCommonMultiple liefert die Position der ersten Zahl in der Liste, die ein gemeinsamer Teiler von a und b ist.
// Falls keine solche Zahl gefunden wird, wird -1 zurückgegeben.
//
// Anmerkung: Wie bisher definieren wir die Funktion für das Suchkriterium direkt hier
// und verwenden dann die Funktion FindFirst, um die Suche durchzuführen.
//
// Es gibt aber noch einen Unterschied: Die Funktion FindFirstCommonMultiple erwartet
// zwei Parameter a und b. Diese Parameter werden hier von der Funktion iscommonmultiple
// *eingefangen*: D.h. sie werden von der Funktion iscommonmultiple verwendet, obwohl sie
// nicht als Parameter an die Funktion iscommonmultiple übergeben werden.
// Die Funktion iscommonmultiple ist eine *Closure*.
func FindFirstCommonMultiple(list []int, a, b int) int {
	iscommonmultiple := func(x int) bool {
		return x%a == 0 && x%b == 0
	}
	return FindFirst(list, iscommonmultiple)
}
