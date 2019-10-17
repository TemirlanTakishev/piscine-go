package piscine

func UltimateDivMod(a *int, b *int) {

	*a = *a / *b
	*b = *b%*a - 1

}
