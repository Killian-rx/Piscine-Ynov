package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 100 {
		return 0
	}
	res := 1
	for i := 1; i <= nb; i++ {
		res = res * i
	}
	return res
}
