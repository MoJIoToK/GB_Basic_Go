package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	filePth := os.Args[1]

	// filePth := "/home/robotomize/1.txt.txt"

	var fileName, fileExt string

	//filePthIndexName индекс последнего слеша в пути к файлу.
	filePthIndexName := strings.LastIndex(filePth, "/")

	//filePthIndexExt индекс последней точки в пути к файлу.
	filePthIndexExt := strings.LastIndex(filePth, ".")
	if filePthIndexExt != -1 {
		fileExt = filePth[filePthIndexExt+1:]
		fileName = filePth[filePthIndexName+1 : filePthIndexExt]
	} else {
		fileExt = ""
		fileName = filePth[filePthIndexName+1:]
	}

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
