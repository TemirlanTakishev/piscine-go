package piscine

func NRune(s string, n int) rune {

	a := []rune(s)

	return a[n] - 31
}
