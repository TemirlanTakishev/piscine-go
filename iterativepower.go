package piscine

func IterativePower(nb int, power int) int {

	if nb > power {
		return nb

		c := power
		b := 1
		for i := 1; i <= nb; i++ {

			b = b * c

		}
		return b
	}
	return nb
}
