package piscine

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	}

	return nb * IterativePower(nb, power-1)
}
