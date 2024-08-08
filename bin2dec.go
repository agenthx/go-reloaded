package piscine

import (
	"math"
	"strconv"
)

func Bin2Dec(binS string) string {
	dec:=0 //decimal
	re:=0 //remainder
	c:=0.0 //counter

	//convert bin string to bin int
	binI, _:= strconv.Atoi(binS)
	//converts the binary number to decimal int
	for binI > 0 {
		re = binI % 10
		binI /= 10
		dec += re * int(math.Pow(2, c))
		c++
	}
	//convert decimal int to decimal string adn return it
	return strconv.Itoa(dec)
}
