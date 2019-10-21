package piscine

import "fmt"

func IterativeFactorial(nb int) int {

	b := 1
	for nb := 1; nb <= 4; nb++ {

		b = b * nb
	}
	return b
}

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
