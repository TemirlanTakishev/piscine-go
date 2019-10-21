package piscine

func IterativePower(nb int, power int) int {

	if power == 0 {
		return 1
	}

	if nb == 1 {
		return 1
	}

	if nb > 1 {
		return nb * IterativePower(nb, power-1)
	}
	return power
}
