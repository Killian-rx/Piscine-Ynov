package piscine

func LastRune(s string) rune {
	l := len(s) - 1
	runes := []rune(s)[l]
	return runes
}
