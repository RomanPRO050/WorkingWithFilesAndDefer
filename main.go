package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	defer executionTime()()
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	readedFile := bufio.NewReader(file)
	newFile, err := os.Create("out.txt")
	defer newFile.Close()
	var numberOfString, bytes int
	for {
		numberOfString += 1
		strNumberOfString := strconv.Itoa(numberOfString) + " "
		line, err := readedFile.ReadString('\n')
		newFile.WriteString(strNumberOfString)
		newFile.WriteString(line)
		bytes += len(line + strNumberOfString)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
	}
	fmt.Printf("Записано %v строк и %v байт", numberOfString, bytes)
}

func executionTime() func() {
	start := time.Now()
	return func() {
		duration := time.Since(start)
		fmt.Println("\nВремя работы программы", duration.Microseconds())
	}
}
