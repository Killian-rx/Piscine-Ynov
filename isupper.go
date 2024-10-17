package piscine

func IsUpper(s string) bool {
	var amount bool
	for i := 0; i <= len(s)-1; i++ {
		if 65 <= []byte(s)[i] && []byte(s)[i] <= 90 {
			amount = true
			continue
		} else {
			amount = false
		}
	}
	return amount
}
