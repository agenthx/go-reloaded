package piscine

import "fmt"

func Punct(sentence string) string {
	result := ""
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' && i < len(sentence)-1 && Ifp(string(sentence[i+1])) {
			for j := i + 1; j < len(sentence); j++ {
				if Ifp(string(sentence[j])) {
					result += string(sentence[j])
				} else if sentence[j] != ' ' {
					break
				}
				i = j
			}
			result += " "
			fmt.Println(result)
		} else {
			result += string(sentence[i])
		}
	}
	return result
}

func Ifp(l string) bool {
	p := ".,!?:;"

	for _, pp := range p {
		if l == string(pp) {
			return true
		}
	}
	return false
}
