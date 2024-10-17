package piscine

func IsLower(s string) bool {
	for i := 0; i <= len(s)-1; i++ {
		if 97 <= []byte(s)[i] && []byte(s)[i] <= 122 {
			continue
		} else {
			return false
		}
	}
	return true
}
