package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Введите предложение")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	//text := "hello world"
	textMap := make(map[string]int)

	text = strings.Trim(text, " ")
	text = strings.Replace(text, " ", "", -1)

	for _, val := range text {
		val = unicode.ToLower(val)

		_, exist := textMap[string(val)]
		if exist {
			textMap[string(val)]++
		} else {
			textMap[string(val)] = 1
		}
	}

	for key, val := range textMap {
		res := float32(val) / float32(len(textMap))
		fmt.Printf("%s - %d %.2f \n", key, val, res)

	}

}
