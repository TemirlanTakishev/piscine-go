package piscine

func IsPrintable(str string) bool {

	p := []rune(str)
	len := 0

	for range str {
		len++
	}
	for i := 0; i < len; i++ {

		if p[i] >= 'A' && p[i] <= 'Z' {
		} else {
			return true
		}
	}

	return false
}
