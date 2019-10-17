package piscine

func UltimateDivMod(a *int, b *int) {
 
	var g int = *a
	*a= *a / *b

    *b = g % *b

}
