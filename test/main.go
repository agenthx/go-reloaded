package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var arr [4]string
	arr[0] = "lol"
	arr[1] = "hahah"
	arr[2] = "kakak"
	arr[3] = "(up, 2)"
	result := "lol "
	for k, word := range arr {
		if strings.HasPrefix(word, "(up,") {
			for i, ch := range word {
				if ch == ',' {
					num, _ := strconv.Atoi(string(word[i+2]))
					for j := num; j > 0; j-- {
						result += strings.ToUpper(string(arr[k-num])) + " "
						num--
					}
				}
			}
		}
	}
	fmt.Println(result)
}
