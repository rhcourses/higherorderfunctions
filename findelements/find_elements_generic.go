package findelements

// ListContains liefert true, falls die Liste list den Wert x enthält.
// Diese Funktion ist generisch, d.h. sie definiert einen *Typparameter* `T` und kann
// für beliebige Typen `T` verwendet werden.
func ListContains[T comparable](list []T, x T) bool {
	for _, v := range list {
		if v == x {
			return true
		}
	}
	return false
}
