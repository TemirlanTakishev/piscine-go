package piscine

func MakeRange(min, max int) []int {
	var b []int

	if min >= max {
		return b
	}

	a := make([]int, max-min)

	for i := range a {
		a[i] = i + min
	}
	return a
}
