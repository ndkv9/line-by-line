package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// open the file with os
	file, err := os.OpenFile("./test.txt", os.O_RDONLY, 0666)
	// handle error and close the program if it happens
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // close the file before exit

	// create new scanner
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
