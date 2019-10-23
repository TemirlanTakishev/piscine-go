package piscine

func ToUpper(s string) string {

	c := []rune(s)

	for i, a := range s {

		if a >= 'a' && a <= 'z' {

			c[i] = a - 32

		}
	}
	return string(c)
}
