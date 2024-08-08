package piscine

import (
	"math"
	"strconv"
)

func Hex2Dec(hex string) string {
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
	//convert dec int to string
	return strconv.Itoa(dec)
}
