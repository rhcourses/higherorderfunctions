package findelements

// FindFirstEven liefert die Position der ersten gerade Zahl in der Liste.
// Falls keine gerade Zahl gefunden wird, wird -1 zurückgegeben.
func FindFirstEven(list []int) int {
	for i, v := range list {
		if v%2 == 0 {
			return i
		}
	}
	return -1
}

// FindFirstOdd liefert die Position der ersten ungeraden Zahl in der Liste.
// Falls keine ungerade Zahl gefunden wird, wird -1 zurückgegeben.
func FindFirstOdd(list []int) int {
	for i, v := range list {
		if v%2 == 1 {
			return i
		}
	}
	return -1
}

// FindFirstGreater3 liefert die Position der ersten Zahl in der Liste, die größer als 3 ist.
// Falls keine Zahl größer als 3 gefunden wird, wird -1 zurückgegeben.
func FindFirstGreater3(list []int) int {
	for i, v := range list {
		if v > 3 {
			return i
		}
	}
	return -1
}
