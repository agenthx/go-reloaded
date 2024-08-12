package main

import (
	"fmt"
	piscine "piscine/functions"
)

func main() {
	var items [7]string
	items[0] = "Okay"
	items[1] = "(wow)"
	items[2] = "lets"
	items[3] = "o"
	items[4] = "hAJDFAKLFa"
	items[5] = "(low,3)"
	items[6] = "10"

	// sent:="okay then lets see this (up,3)"
	// arr:= strings.Split(sent, " ")
	// result:=""
	// for k, word := range arr {
	// 	if strings.HasPrefix(word, "(up,") {
	// 		for i, ch := range word {
	// 			if ch == ',' {
	// 				num, _ := strconv.Atoi(string(word[i+]))
	// 				for j := num; j > 0; j-- {
	// 					result += strings.ToUpper(string(arr[k-num])) + " "
	// 					num--
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	// fmt.Println(result)
	fmt.Println(piscine.UpLowCap(items))
}
