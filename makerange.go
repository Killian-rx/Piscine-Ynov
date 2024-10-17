package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return nil
	}
	size := max - min
	res := make([]int, size)
	for i := 0; i < size; i++ {
		min = min + 1
		res[i] = min - 1
	}
	return res
}
