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
	var itemsS []string
	for scanner.Scan() {
		line := scanner.Text()
		lineItems := strings.Fields(line)
		itemsS = append(itemsS, lineItems...)
	}
	var sent string
	for _, word := range itemsS {
		sent += string(word) + " "
	}
	//fix the editing functions with regex
	reU:=regexp.MustCompile(`(?i)\(\s*up\s*\)`)
	sent = reU.ReplaceAllString(sent, "(up)")//up
	reL:=regexp.MustCompile(`(?i)\(\s*low\s*\)`)
	sent = reL.ReplaceAllString(sent, "(low)")//low
	reC:=regexp.MustCompile(`(?i)\(\s*cap\s*\)`)
	sent = reC.ReplaceAllString(sent, "(cap)")//cap
	reH:=regexp.MustCompile(`(?i)\(\s*hex\s*\)`)
	sent = reH.ReplaceAllString(sent, "(hex)")//hex
	reB:=regexp.MustCompile(`(?i)\(\s*bin\s*\)`)
	sent = reB.ReplaceAllString(sent, "(bin)")//bin
	reUN:=regexp.MustCompile(`(?i)\(\s*up\s*,\s*(\d*)\s*\)`)
	sent = reUN.ReplaceAllString(sent, "(up,$1)") //up w num
	reLN:=regexp.MustCompile(`(?i)\(\s*low\s*,\s*(\d*)\s*\)`)
	sent = reLN.ReplaceAllString(sent, "(low,$1)") //low w num
	reCN:=regexp.MustCompile(`(?i)\(\s*cap\s*,\s*(\d*)\s*\)`)
	sent = reCN.ReplaceAllString(sent, "(cap,$1)") //cap w num

	//turn the sent to an array of string
	items:=strings.Fields(sent)

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
	// UpLowCap with numbers
	sentence := piscine.UpLowCap(items)
	// remove the (hex), etc..
	regex := regexp.MustCompile(`\(bin\)|\(hex\)|\(up\)|\(low\)|\(cap\)|\(up,\d+\)|\(low,\d+\)|\(cap,\d+\)`)
	sentence = regex.ReplaceAllLiteralString(sentence, "")
	// remove extra spaces caused by regex
	sentence = piscine.Spaces(sentence)
	// fix punctuations
	sentence = piscine.Punct(sentence)
	// fix vowels
	sentence = piscine.Vowels(sentence)
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
