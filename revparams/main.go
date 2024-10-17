package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i > 0; i-- {
		temp := os.Args[i]
		runes := []rune(temp)
		for _, r := range runes {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n') // Nouvelle ligne à la fin
	}
}
