package main

import "fmt"

func PointOne(n *int) {
	*n = 1

}
func main() {
	n := 5
	PointOne(&n)
	fmt.Println(n)
}
