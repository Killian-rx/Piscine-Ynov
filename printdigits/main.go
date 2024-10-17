package main

import "github.com/01-edu/z01"

var alphab = [10]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func main() {
	for i := 0; i < 10; i++ {
		z01.PrintRune(alphab[i])
		if alphab[i] == '9' {
			z01.PrintRune('\n')
		}
	}
}
