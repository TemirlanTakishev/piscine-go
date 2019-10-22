package piscine

func NRune(s string, n int) rune {
	result := []rune(s)

	for index, value := range result {
		if index == n-1 {
			return value
		}
	}
	return 0
}
