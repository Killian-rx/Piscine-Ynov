package piscine

func Atoi(s string) int {
	var res int
	var sign int = 1
	for j, i := range s {
		if j == 0 && i == '-' {
			sign = -1
		} else if j == 0 && (i == '+' || i == ' ') {
		} else if i >= '0' && i <= '9' {
			res = res*10 + int(i-'0')
		} else {
			return 0
		}
	}
	return res * sign
}
