package piscine

func ToLower(s string) string {

	c := []rune(s)

	for i, a := range s {

		if a >= 'A' && a <= 'Z' {

			c[i] = a + 32

		}
	}
	return string(c)
}
