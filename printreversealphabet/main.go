package main

import "github.com/01-edu/z01"

var alphab = [26]rune{'z', 'y', 'x', 'w', 'v', 'u', 't', 's', 'r', 'q', 'p', 'o', 'n', 'm', 'l', 'k', 'j', 'i', 'h', 'g', 'f', 'e', 'd', 'c', 'b', 'a'}

func main() {
	for i := 0; i < 26; i++ {
		z01.PrintRune(alphab[i])
		if alphab[i] == 'a' {
			z01.PrintRune('\n')
		}
	}
}
