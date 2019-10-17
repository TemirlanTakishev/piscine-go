package piscine

func StrRev(s string) string {
	Barbara := []byte(s)
	Glen := 0

	for range Barbara {
		Glen++
	}

	for i, c := range []byte(s) {
		Barbara[Glen-i-1] = c

	}

	return string(Barbara)
}
