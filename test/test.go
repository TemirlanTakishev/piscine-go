package main

import (
	"fmt"

	piscine ".."
)

func main() {
	fmt.Println(piscine.Index("Hello!", "Hello!rus"))
	fmt.Println(piscine.Index("Salut!", "alu"))
	fmt.Println(piscine.Index("Ola!", "hOl"))
}
