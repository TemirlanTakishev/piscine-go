package piscine

func IsLower(str string) bool {

	Ran := []rune(str)
	len := 0

	for range str {
		len++
	}
	for i := 0; i < len; i++ {

		if Ran[i] >= 'a' && Ran[i] <= 'z' {
		} else {
			return false
		}
	}
	return true
}
