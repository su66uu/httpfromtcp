package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func fClose(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatal("Error closing the file", "error", err)
	}
}

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("Error", "error", err)
	}

	defer fClose(f)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("read: %s\n", line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading the file", err)
	}
}
