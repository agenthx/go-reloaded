package piscine

import (
	"strconv"
	"strings"
)

func UpLowCap(items [7]string) string {
	var num int   // the number given
	var place int // place of function
	what := 0
	for i, word := range items {
		if strings.HasPrefix(word, "(up,") { // we found the up with number
			for _, ch := range word {
				if ch > '0' && ch < '9' {
					num, _ = strconv.Atoi(string(ch)) // we extracted the number
				}
			}
		}
	}
	// switch
	switch what {
	case 1:
		// up
		items[i-j] = strings.ToUpper(items[i-j])
	case 2:
		// low
		items[i-j] = strings.ToLower(items[i-j])

	case 3:
		// cap
		items[i-j] = strings.Title(items[i-j])
	default:
		// do nothing
	}

	// store the updated items in a string
	sent := ""
	for _, word := range items {
		sent += word + " "
	}
	return sent
}
