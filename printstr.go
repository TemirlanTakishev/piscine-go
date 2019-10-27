package main

import "fmt"

func StrLen(str string) int {

	len := 0

	for range str {
		len++
	}
	return len
}

func main() {
	str := "Hello World!"
	nb := StrLen(str)
	fmt.Println(nb)
}
