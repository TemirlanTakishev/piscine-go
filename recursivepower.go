package piscine

func RecursivePower(nb int, power int) int {

	if power == 0 {
		return 1
	} else if nb < 1 {
		return 0
	}
	return nb * RecursivePower(nb, power-1)
}
