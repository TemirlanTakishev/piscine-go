package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	rik := os.Args[1:]        // мы не создаем массив т,к если в задании пишут аргументы
	for _, arg := range rik { // нам не нужны индексы мы проходимся по буквам для того чтобы пройтись по словам
		for _, letter := range arg { // вывести слова по символьно

			z01.PrintRune(letter) // печатай нам символы
		}
		z01.PrintRune(10)
	}

}
