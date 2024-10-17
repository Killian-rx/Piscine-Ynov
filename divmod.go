package piscine

func DivMod(a int, b int, div *int, mod *int) {
	res := a / b
	rest := a % b
	*div = res
	*mod = rest
}
