package piscine

func UltimateDivMod(a *int, b *int) {
	res := *a / *b
	rest := *a % *b
	*a = res
	*b = rest
}
