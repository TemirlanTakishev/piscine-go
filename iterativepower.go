package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	c := power
	b := 1
	for i := 1; i <= nb; i++ {

		b = b * c

	}
	return b
}
