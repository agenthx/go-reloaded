// package piscine

// import (
// 	"fmt"
// 	"os"
// 	"regexp"
// 	"strconv"
// 	"strings"
// )

// func InputConversions() string {
// 	txt, err := os.ReadFile(os.Args[1])
// 	if err != nil {
// 		fmt.Print(err)
// 	}

// 	var combine string = string(txt)

// 	//recombining the different cases if there are spaces in-between and adjusting for case-insensitivity
// 	reH := regexp.MustCompile(`(?i)\(+\s*h\s*e\s*x\s*\)+`)
// 	combine = reH.ReplaceAllString(combine, " (hex) ")
// 	reB := regexp.MustCompile(`(?i)\(+\s*b\s*i\s*n\s*\)+`)
// 	combine = reB.ReplaceAllString(combine, " (bin) ")
// 	reL := regexp.MustCompile(`(?i)\(+\s*l\s*o\s*w\s*\)+`)
// 	combine = reL.ReplaceAllString(combine, " (low) ")
// 	reU := regexp.MustCompile(`(?i)\(+\s*u\s*p\s*\)+`)
// 	combine = reU.ReplaceAllString(combine, " (up) ")
// 	reC := regexp.MustCompile(`(?i)\(+\s*c\s*a\s*p\s*\)+`)
// 	combine = reC.ReplaceAllString(combine, " (cap) ")
// 	reCNum := regexp.MustCompile(`(?i)\(+\s*c\s*a\s*p\s*,\s*(\d+)\s*\)+`)
// 	combine = reCNum.ReplaceAllString(combine, " (cap,$1) ")
// 	reUNum := regexp.MustCompile(`(?i)\(+\s*u\s*p\s*,\s*(\d+)\s*\)+`)
// 	combine = reUNum.ReplaceAllString(combine, " (up,$1) ")
// 	reLNum := regexp.MustCompile(`(?i)\(+\s*l\s*o\s*w\s*,\s*(\d+)\s*\)+`)
// 	combine = reLNum.ReplaceAllString(combine, " (low,$1) ")

// 	//turns the string into a splice of strings
// 	str := strings.Fields(combine)

// 	for i := 0; i < len(str); i++ {
// 		//converts from hex to dec
// 		if reH.MatchString(str[i]) {
// 			hexDecValue, err := strconv.ParseInt(str[i-1], 16, 64) //converting hex value to dec
// 			if err != nil {
// 				fmt.Print(err)
// 			}
// 			result := strconv.Itoa(int(hexDecValue))
// 			str[i-1] = result
// 			str = append(str[:i], str[i+1:]...) //code for skipping the current substring i
// 			i--

// 			//converts from bin to dec
// 		} else if reB.MatchString(str[i]) {
// 			binDecValue, err := strconv.ParseInt(str[i-1], 2, 64)
// 			if err != nil {
// 				fmt.Print(err)
// 			}
// 			result := strconv.Itoa(int(binDecValue))
// 			str[i-1] = result
// 			str = append(str[:i], str[i+1:]...)
// 			i--

// 			//converts from lower to upper case
// 		} else if reU.MatchString(str[i]) {
// 			str[i-1] = strings.ToUpper(str[i-1])
// 			str = append(str[:i], str[i+1:]...)
// 			i--

// 			//converts from upper to lower case
// 		} else if reL.MatchString(str[i]) {
// 			str[i-1] = strings.ToLower(str[i-1])
// 			str = append(str[:i], str[i+1:]...)
// 			i--

// 			//converts from uncapitalized to cap
// 		} else if reC.MatchString(str[i]) {
// 			str[i-1] = Capitalize(str[i-1])
// 			str = append(str[:i], str[i+1:]...)
// 			i--

// 			//matches for (up, Y), where Y is any digit
// 		} else if reUNum.MatchString(str[i]) {
// 			match := reUNum.FindStringSubmatch(str[i]) //provides slice with index[0] being the match and index[1] being captured group
// 			//  ((?i)\(+\s*c\s*a\s*p\s*,\s*(\d+)\s*\)+)
// 			//match[0] =  
// 			if len(match) > 1 {
// 				digitstr := match[1]
// 				digit, err := strconv.Atoi(digitstr)
// 				if err != nil {
// 					fmt.Print(err)
// 				}

// 				if digit > i {
// 					fmt.Print("Error: digit must be of proper length")
// 					break
// 				}

// 				for j := 0; j < digit; j++ {
// 					str[i-digit+j] = strings.ToUpper(str[i-digit+j])
// 				}
// 				str = append(str[:i], str[i+1:]...)
// 				i -= digit
// 			}

// 			//matches for (low, Y)
// 		} else if reLNum.MatchString(str[i]) {
// 			match := reLNum.FindStringSubmatch(str[i])

// 			if len(match) > 1 {
// 				digitstr := match[1]
// 				digit, err := strconv.Atoi(digitstr)
// 				if err != nil {
// 					fmt.Print(err)
// 				}

// 				if digit > i {
// 					fmt.Print("Error: digit must be of proper length")
// 					break
// 				}

// 				for j := 0; j < digit; j++ {
// 					str[i-digit+j] = strings.ToLower(str[i-digit+j])
// 				}
// 				str = append(str[:i], str[i+1:]...)
// 				i -= digit
// 			}

// 			//matches for (cap, Y)
// 		} else if reCNum.MatchString(str[i]) {
// 			match := reCNum.FindStringSubmatch(str[i])

// 			if len(match) > 1 {
// 				digitstr := match[1]
// 				digit, err := strconv.Atoi(digitstr)
// 				if err != nil {
// 					fmt.Print(err)
// 				}

// 				if digit > i {
// 					fmt.Print("Error: digit must be of proper length")
// 					break
// 				}

// 				for j := 0; j < digit; j++ {
// 					str[i-digit+j] = Capitalize(str[i-digit+j])
// 				}
// 				str = append(str[:i], str[i+1:]...)
// 				i -= digit
// 			}
// 		}
// 	}
// 	final := strings.Join(str, " ")
// 	return final
// f }