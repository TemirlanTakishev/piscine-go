package piscine

func RecursivePower(nb int, power int) int {

	if power == 0 {
		return 1
	}

	if nb == 1 {
		return 1
	}

	if nb > 1 {
		return nb * RecursivePower(nb, power-1)
	}
	return power
}
