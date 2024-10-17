package piscine

func Any(f func(string) bool, a []string) bool {
	res := false
	for i := 0; i < len(a); i++ {
		result := f(a[i])
		if result {
			res = true
		}
	}
	return res
}
