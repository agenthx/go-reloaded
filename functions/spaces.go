package piscine

func Spaces(sentence string) string {
	result := ""
	flag := false
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' && flag {
			result += string(sentence[i])
			flag = false
		} else if sentence[i] != ' ' {
			result += string(sentence[i])
			flag = true
		}
	}
	return result
}
