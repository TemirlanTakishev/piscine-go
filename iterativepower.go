package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}

	b := 1
	for i := 1; i <= nb; i++ {

		b = nb * i * i

	}
	return b
}
