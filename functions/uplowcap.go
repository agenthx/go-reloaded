package piscine

import (
	"strconv"
	"strings"
)

func UpLowCap(items []string) string {
	var num int   // the number given
	var place int // place of function
	what := 0     //which function to perform
	for i, word := range items {
		if strings.HasPrefix(word, "(up,") { // we found the up with number
			what = 1
			place = i
			num = ExtractNum(word)
			items = modify(what, place, num, items)
		} else if strings.HasPrefix(word, "(low,") { // we found the low with number
			what = 2
			place = i
			num = ExtractNum(word)
			items = modify(what, place, num, items)
		} else if strings.HasPrefix(word, "(cap,") { // we found the cap with number
			what = 3
			place = i
			num = ExtractNum(word)
			items = modify(what, place, num, items)
		}
	}
	// store the updated items in a string
	sent := ""
	for _, word := range items {
		sent += word + " "
	}
	return sent
}

func ExtractNum(word string) int {
	var num int
	for _, ch := range word {
		if ch > '0' && ch < '9' {
			num, _ = strconv.Atoi(string(ch)) // we extracted the number
		}
	}
	return num
}

func modify(what, place, num int, items []string) []string {
	// switch
	switch what {
	case 1:
		// up
		for j := num; j > 0; j-- {
			items[place-j] = strings.ToUpper(items[place-j])
		}
	case 2:
		// low
		for j := num; j > 0; j-- {
			items[place-j] = strings.ToLower(items[place-j])
		}
	case 3:
		// cap
		for j := num; j > 0; j-- {
			items[place-j] = strings.Title(items[place-j])
		}
	}
	return items
}
