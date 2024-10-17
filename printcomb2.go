package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			k := j + 1
			for l := i; l <= '9'; l++ {
				for ; k <= '9'; k++ {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(l)
					z01.PrintRune(k)
					if i < '9' || j < '8' || l < '9' || k < '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				k = '0'
			}
		}
	}
	z01.PrintRune('\n')
}
