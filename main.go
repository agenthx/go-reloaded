package main

import (
	"fmt"
	"os"
	piscine "piscine/up"
)

func main() {
	input := "sample.txt"
	output := "result.txt"

	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()

	data := make([]byte, 34)
	sentence := ""

	for {
		i, err := file.Read(data)
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Error:", err)
		}
		if i == 0 {
			break
		}
		sentence += string(data[:i])
	}

	sentence , err = piscine.Up(sentence)
	if err != nil {
		fmt.Println(err)
		return
	}

	File, err := os.Create(output)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer File.Close()

	_, err = File.WriteString(sentence)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Ready")

}
