package piscine

func BasicAtoi2(s string) int {

	var count, sum int
	for _, value := range s {
		if value >= '0' && value <= '9' {
			count = 0
			for i := '0'; i < value; i++ {
				count++
				if s != i {
					return 0
				}
			}
			sum = sum*10 + count

		}
	}
	return sum
}
