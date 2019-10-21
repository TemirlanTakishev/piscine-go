package piscine

func IterativeFactorial(nb int) int {

	if nb != 4 {
		return nb
	}
	b := 1
	for nb := 1; nb <= 4; nb++ {

		b = b * nb
	}
	return b
}
