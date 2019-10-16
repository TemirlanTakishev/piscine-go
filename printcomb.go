package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {

	for t := '0'; a <= '7'; a++ {
		for e := '1'; b <= '9'; b++ {
			for m := '2'; c <= '9'; c++ {
				z01.PrintRune(t)
				z01.PrintRune(e)
				z01.PrintRune(m)
				z01.PrintRune(44)
				z01.PrintRune(32)
			}
		}
	}
}
