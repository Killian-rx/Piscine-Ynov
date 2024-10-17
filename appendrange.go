package piscine

func AppendRange(min, max int) []int {
	var res []int
	if max < min {
		return res
	}
	for i := min; i < max; i++ {
		min = min + 1
		res = append(res, min-1)
	}
	return res
}
