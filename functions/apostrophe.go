package piscine

func Apostrophe(sent string) string {
	result := ""
	flag := true
	for i := 0; i < len(sent); i++ {
		if sent[i] == '\'' && flag {
			result += string(sent[i])
			flag = false
			for j := i + 1; j < len(sent); j++ {
				if sent[j] != ' ' && sent[j] != '\'' {
					result += string(sent[j])
				} else if sent[j] == '\'' {
					result += string(sent[j])
					i = j
					j = len(sent)
				}
			}
		} else {
			result += string(sent[i])
		}
	}
	return result
}
