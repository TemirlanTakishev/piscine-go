package piscine

func IterativePower(nb int, power int) int {

	if power > 0 {
		return 0
	}
	b := 1
	for i := 0; i < power; i++ {

		b = b * i

	}

	return b
}
