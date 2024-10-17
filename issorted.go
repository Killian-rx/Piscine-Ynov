package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) <= 1 {
		return true
	}
	sortc := true
	sortd := true
	for i := 0; i < len(a)-1; i++ {
		result := f(a[i], a[i+1])
		if result > 0 {
			sortc = false
		} else if result < 0 {
			sortd = false
		}
	}
	return sortc || sortd
}
