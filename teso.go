package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func ABO() {

	Per := []rune(os.Args[0])

	for _, letter := range Per {

		z01.PrintRune(letter)
	}
	z01.PrintRune(10)
}
