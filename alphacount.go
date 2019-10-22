package piscine

func AlphaCount(str string) int {

	a := 0

	for _, i := range str {

		if (i >= 'a' && i <= 'z') || (i > 'A' && i <= 'Z') {
			a++
		}

	}

	return a
}
