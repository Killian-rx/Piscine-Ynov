package piscine

func IsAlpha(s string) bool {
	for i := 0; i <= len(s)-1; i++ {
		if (97 <= []byte(s)[i] && []byte(s)[i] <= 122) || (65 <= []byte(s)[i] && []byte(s)[i] <= 90) || (48 <= []byte(s)[i] && []byte(s)[i] <= 57) {
			continue
		} else {
			return false
		}
	}
	return true
}
