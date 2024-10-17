package piscine

func Map(f func(int) bool, a []int) []bool {
	var res []bool
	for i := 0; i < len(a); i++ {
		result := f(a[i])
		res = append(res, result)
	}
	return res
}
