package piscine

func TrimAtoi(s string) int {
	var res int
	var sign int = 1
	lenght := 0
	for j, i := range s {
		if i >= '0' && i <= '9' {
			res = res*10 + int(i-'0')
			lenght = j + 1
		} else if i == '-' && lenght == 0 {
			sign = -1
		} else if i == '+' {
			sign = 1
		}
	}
	return res * sign
}
