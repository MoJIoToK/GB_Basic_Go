package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Введите предложение: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	var wordLine, symbolNumber, bytesNumber int

	bytesNumber = len(text)

	strWithoutSpace := strings.Trim(text, " ")
	fmt.Println(strWithoutSpace)

	for _, v := range strWithoutSpace {

		if unicode.IsSpace(v) {
			wordLine++
		}

		if unicode.IsLetter(v) {
			symbolNumber++
		}

	}

	fmt.Printf("Количество слов: %d ", wordLine-1)
	fmt.Printf("Количество букв: %d ", symbolNumber)
	fmt.Printf("Количество байт: %d ", bytesNumber)
}
