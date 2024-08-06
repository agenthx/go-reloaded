package piscine

import (
	"fmt"
	"regexp"
	"strings"
)

func Up(sentence string) (string , error){
	arr := strings.Split(sentence, " ")
	match := true
	for i, word := range arr {
		if word == "(up)" && i == 0{
			return "" , fmt.Errorf("nothing before up")
		} else if word == "(low)" && i == 0{
			return "" , fmt.Errorf("nothing before low")
		}
		if word == "(up)" {
			arr[i-1] = strings.ToUpper(arr[i-1])
			match = false
		}
	}

	if !match {
		sentence = ""
		for i, word := range arr {
			sentence += word
			if i < len(arr)-1 {
				sentence += " "
			}
		}
	}

	rule := regexp.MustCompile(`\(up\)`)
	sentence = rule.ReplaceAllLiteralString(sentence, "")

	return sentence , nil
}
