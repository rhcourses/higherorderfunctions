package filter

// listsequal prüft, ob zwei Listen gleich sind.
// Auch dies ist eine generische Funktion.
func listsequal[T comparable](l1, l2 []T) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i, v := range l1 {
		if v != l2[i] {
			return false
		}
	}
	return true
}
