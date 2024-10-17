package piscine

func Sqrt(nb int) int {
	if nb < 0 || nb > 100 {
		return 0
	}
	if nb == 1 {
		return 1
	}
	res := 0
	for i := 0; i < nb; i++ {
		if nb == i*i {
			res = i
		}
	}
	return res
}
