package piscine

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if nb < 0 {
		return 0
	}

	return nb * IterativePower(nb, power-1)
}
