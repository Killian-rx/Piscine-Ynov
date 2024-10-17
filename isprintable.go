package piscine

func IsPrintable(s string) bool {
	for i := 0; i <= len(s)-1; i++ {
		if 32 <= []byte(s)[i] && []byte(s)[i] <= 126 {
			continue
		} else {
			return false
		}
	}
	return true
}
