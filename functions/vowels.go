package piscine

func Vowels(sentence string) string {
	result := ""

	for i, ch := range sentence {
		if (ch == 'a' || ch == 'A') && string(sentence[i+1]) == " " && vow(string(sentence[i+2])) {
			result += string(ch) + "n"
		} else {
			result += string(ch)
		}
	}
	return result
}

func vow(word string) bool {
	letter := word[0]
	vowels := "aeiou"

	for _, ch := range vowels {
		if letter == byte(ch) {
			return true
		}
	}
	return false
}
