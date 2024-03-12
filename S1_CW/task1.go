package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func isPrime(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i < sqrt; i++ {
		//fmt.Println(i)
		if n%i == 0 {
			//fmt.Println("false", i)
			return false
		}
	}
	return true
}

func main() {

	var value int
	fmt.Println("Введите число")
	if _, err := fmt.Fscan(os.Stdin, &value); err != nil {
		log.Fatal(err)
	}

	if isPrime(value) {
		fmt.Println("Число", value, "простое")
	} else {
		fmt.Println("Число", value, "не простое")
	}
}
