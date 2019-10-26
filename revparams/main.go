package main 

import (
	"os"
)

func main(){
	per := []rune(os.Args[1:]) // создаем масив рун из библиотеки os Args

	for i :range per {
		for i,  := range per [i]{
			
			z01.PrintRune()
		}
	} 

}