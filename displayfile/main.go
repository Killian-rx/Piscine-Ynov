package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Printf("Too many arguments")
		fmt.Printf("\n")
		return
	} else if len(os.Args) < 2 {
		fmt.Printf("File name missing")
		fmt.Printf("\n")
		return
	}
	filename := os.Args[1]
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file")
		return
	}
	fmt.Print(string(file))
}
