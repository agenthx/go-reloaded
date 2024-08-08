package piscine

import (
	"fmt"
	"regexp"
	"strings"
)

func UpLowCap(sentence string) (string, error) {
	arr := strings.Split(sentence, " ")
	flag := true
	for i, word := range arr {
		if word == "(up)" && i == 0 {
			return "", fmt.Errorf("nothing before up")
		} else if word == "(low)" && i == 0 {
			return "", fmt.Errorf("nothing before low")
		} else if word == "(cap)" && i == 0 {
			return "", fmt.Errorf("nothin before cap")
		}
		if word == "(up)" {
			arr[i-1] = strings.ToUpper(arr[i-1])
			flag = false
		} else if word == "(low)" {
			arr[i-1] = strings.ToLower(arr[i-1])
			flag = false
		} else if word == "(cap)" {
			arr[i-1] = strings.Title(arr[i-1])
			flag = false
		}
	}

	//go here if you find up or low
	if !flag {
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

	rule1 := regexp.MustCompile(`\(low\)`)
	sentence = rule1.ReplaceAllLiteralString(sentence, "")

	rule2 := regexp.MustCompile(`\(cap\)`)
	sentence = rule2.ReplaceAllLiteralString(sentence, "")

	return sentence, nil
}
