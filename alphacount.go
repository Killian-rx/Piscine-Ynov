package piscine

func AlphaCount(s string) int {
	amount := 0
	for i := 0; i <= len(s)-1; i++ {
		if (97 <= []byte(s)[i] && []byte(s)[i] <= 122) || (65 <= []byte(s)[i] && []byte(s)[i] <= 90) {
			amount++
		}
	}
	return amount
}
