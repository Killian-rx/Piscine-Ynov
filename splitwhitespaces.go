package piscine

func SplitWhiteSpaces(s string) []string {
	var res []string
	var word string
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\t' && s[i] != '\n' {
			word += string(s[i])
		} else {
			if word != "" {
				res = append(res, word)
			}
			word = ""
		}
	}

	if word != "" {
		res = append(res, word)
	}
	return res
}
