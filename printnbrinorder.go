package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	length := 1
	temp := n
	for temp >= 10 {
		temp /= 10
		length++
	}
	digits := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		digit := n % 10
		digits[i] = digit
		n /= 10
	}
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if digits[i] > digits[j] {
				digits[i], digits[j] = digits[j], digits[i]
			}
		}
	}
	for _, digit := range digits {
		z01.PrintRune(rune(digit) + '0')
	}
}
