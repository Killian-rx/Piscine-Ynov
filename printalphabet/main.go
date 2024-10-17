package main

import "github.com/01-edu/z01"

var alphab = [26]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func main() {
	for i := 0; i < 26; i++ {
		z01.PrintRune(alphab[i])
		if alphab[i] == 'z' {
			z01.PrintRune('\n')
		}
	}
}
