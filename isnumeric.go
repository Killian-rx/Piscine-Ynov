package piscine

func IsNumeric(s string) bool {
	for i := 0; i <= len(s)-1; i++ {
		if 48 <= []byte(s)[i] && []byte(s)[i] <= 57 {
			continue
		} else {
			return false
		}
	}
	return true
}
