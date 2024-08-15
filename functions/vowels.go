package piscine

import (
	"fmt"
	"regexp"
	"strings"
)

func Vowels(sentence string) string {
	result := ""
	arr := strings.Fields(sentence)

	for i, ch := range arr {
		if ch == "a" && vow(string(arr[i+1])) {
			arr[i] = "an"
		} else if ch == "A" && vow(string(arr[i+1])) {
			arr[i] = "AN"
		} else if ch == "an" && !vow(string(arr[i+1])) {
			arr[i] = "a"
		} else if ch == "AN" && !vow(string(arr[i+1])) {
			arr[i] = "A"
		}
	}
	for _, word := range arr {
		result += word + " "
	}
	return result
}

func vow(word string) bool {
	fmt.Println(word)
	reV := regexp.MustCompile(`(?i)\b[aeiou]`)
	return reV.MatchString(word)
}