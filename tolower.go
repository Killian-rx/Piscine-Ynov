package piscine

func ToLower(s string) string {
	res := []byte(s)
	for i := 0; i <= len(s)-1; i++ {
		if 'A' <= s[i] && s[i] <= 'Z' {
			res[i] = s[i] + 32
		}
	}
	return string(res)
}
