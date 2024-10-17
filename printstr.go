package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	fin := []rune(s)
	for i := range fin {
		z01.PrintRune(fin[i])
	}
}
