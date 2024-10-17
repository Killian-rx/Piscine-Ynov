package piscine

func CountIf(f func(string) bool, tab []string) int {
	nb := 0
	for i := 0; i < len(tab); i++ {
		result := f(tab[i])
		if result {
			nb = nb + 1
		}
	}
	return nb
}
