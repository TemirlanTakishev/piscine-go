package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	len := 0
	for range args {
		len++
	}

	for i := len - 1; i >= 0; i-- {

		param := args[i]

		for _, c := range param {
			z01.PrintRune(c)
		}
		z01.PrintRune(10)
	}

}
