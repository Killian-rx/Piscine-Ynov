package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	res := 1
	for i := 1; i <= power; i++ {
		res = res * nb
	}
	return res
}
