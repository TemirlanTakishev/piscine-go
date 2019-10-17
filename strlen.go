package piscine

func StrLen(str string) int {

	a := 0
	for _, i := range str {
		if i == i {
			a++
		}
	}
	return a
}
