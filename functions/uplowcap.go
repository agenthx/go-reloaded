package piscine

import (
	"strconv"
	"strings"
)

func UpLowCap(items [7]string) string {
	var num int //the number given
	var place int //place of function
	for i, word := range items {
		if strings.HasPrefix(word, "(up,") { //we found the up with number
			for _, ch := range word {
				if ch > '0' && ch < '9' {
					num, _ = strconv.Atoi(string(ch)) //we extracted the number
				}
			}
		}
	}

	//do a case for this
	for j := num; j > 0; j-- {
		items[i-j] = strings.ToUpper(items[i-j])
	}

	//store the updated items in a string
	sent := ""
	for _, word := range items {
		sent += word + " "
	}
	return sent
}
