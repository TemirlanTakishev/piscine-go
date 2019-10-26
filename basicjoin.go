//package piscine

// func BasicJoin(strs []string) string {
// 	a := ""
// 	len := 0
// 	for range strs {
// 		len++
// 	}

// 	for i := 0; i < len; i++ {
// 		a = strs[0] + strs[i] + strs[]
// 	}
// 	return a
// }

package piscine

import "github.com/01-edu/z01"

func PrintComb() {

	for i := '0'; i <= '7'; i++ {
		for j := '0'; j <= '8'; j++ {
			for g := '0'; g <= '9'; g++ {
				if i < j && j < g {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(g)
					z01.PrintRune(32)
					if i != '7' || j != '8' || g != '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}
