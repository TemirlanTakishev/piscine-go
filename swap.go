package main

import "fmt"

func Swap(a *int, b *int) {

	q := *a

	*a = *b
	*b = q
}
func main() {
	a := 0
	b := 1
	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
