package piscine

func Swap(a *int, b *int) {

	q := *a

	*a = *b
	*b = q
}
