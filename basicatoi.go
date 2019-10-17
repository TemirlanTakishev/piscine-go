package piscine

func BasicAtoi(s string) int {
	copy := []rune(s)
	result := 0

	Glen := 0
	for _, c := range copy {
		if c > 57 || c < 48 {
			return 0
		}
		Glen++
	}

	for i := range copy {

		pow := 1
		for j := 0; j < i; j++ {
			pow *= 10
		}

		f := '0'
		n := 0
		for f != (copy[Glen-i-1]) {
			n++
			f++
		}

		result += n * pow

	}

	return result
}
