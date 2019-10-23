package piscine

func LastRune(s string) rune {
	runes := []rune(s)
	len := RuneLength(runes)
	return runes[len-1]
}

func RuneLength(runes []rune) int {
	count := 0
	for range runes {
		count++
	}

	return count
}
