package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:] // Exclut le nom du programme, qui est le premier argument
	for _, arg := range arguments {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n') // Nouvelle ligne à la fin
	}
}
