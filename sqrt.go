package piscine

func Sqrt(nb int) int {

	for i := 0; i <= 9; i++ {
		if i*i == nb {
			return i

		}
	}
	return 0
}
