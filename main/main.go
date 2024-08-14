package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	piscine "piscine/functions"
)

func main() {
	arr:=os.Args //save the args given

	input:=arr[1]
	output:=arr[2]

	file, ferr := os.Open(input) // open the file sample.txt
	if ferr != nil {
		panic(ferr)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // read the contents of the file and make the changes
	var items []string
	for scanner.Scan() {
		line := scanner.Text()
		items = strings.Split(line, " ")
		for i, word := range items {
			if word == "(bin)" {
				items[i-1] = piscine.Bin2Dec(items[i-1])
			} else if word == "(hex)" {
				items[i-1] = piscine.Hex2Dec(items[i-1])
			} else if word == "(up)" {
				items[i-1] = strings.ToUpper(items[i-1])
			} else if word == "(low)" {
				items[i-1] = strings.ToLower(items[i-1])
			} else if word == "(cap)" {
				items[i-1] = strings.Title(items[i-1])
			}
		}
	}
	// UpLowCap with numbers
	sentence := piscine.UpLowCap(items)
	// remove the (hex), etc..
	regex := regexp.MustCompile(`\(bin\)|\(hex\)|\(up\)|\(low\)|\(cap\)|\(up,\d+\)|\(low,\d+\)|\(cap,\d+\)`)
	sentence = regex.ReplaceAllLiteralString(sentence, "")
	// fix vowels
	sentence = piscine.Vowels(sentence)
	// remove extra spaces caused by regex
	sentence = piscine.Spaces(sentence)
	// fix punctuations
	sentence = piscine.Punct(sentence)
	// fix apostrophes
	sentence = piscine.Apostrophe(sentence)
	// create result file
	file, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// write the sentence in the file
	_, werr := file.WriteString(sentence)
	if err != nil {
		panic(werr)
	}
	defer file.Close()

	fmt.Println("READY!!!")
}
