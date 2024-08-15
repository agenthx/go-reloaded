package piscine

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Hex2Dec(hex string) (string, error) {
	hex = strings.ToUpper(hex) //capatalize the hex
	reH := regexp.MustCompile(`[^0-9A-F-]`)
	if reH.MatchString(hex) {
		return "", fmt.Errorf("string '%s' contains non hex values (0-9 A-F)", hex)
	}
	hexI := make([]int, len(hex)) //array hexI to store each hex converted to its int
	c := 0.0                      // counter
	dec := 0                      //decimal

	//converting letters to ints
	for i, ch := range hex {
		if ch >= 'A' && ch <= 'Z' {
			hexI[i] = 10 + int(ch-'A')
		} else {
			hexI[i], _ = strconv.Atoi(string(ch))
		}
	}
	//convert hex to dec
	for i := len(hexI) - 1; i >= 0; i-- {
		dec += hexI[i] * int(math.Pow(16, c))
		c++
	}
	if hexI[0] == 0 {
		return "-" + strconv.Itoa(dec), nil
	}
	//convert dec int to string
	return strconv.Itoa(dec), nil
}