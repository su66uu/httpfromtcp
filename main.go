package main

import (
	"bytes"
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

	str := ""
	for {
		data := make([]byte, 8)
		n, err := f.Read(data)
		if err != nil {
			break
		}

		data = data[:n]
		if i := bytes.IndexByte(data, '\n'); i != -1 {
			str += string(data[:i])
			data = data[i+1:]
			fmt.Printf("read: %s\n", str)
			str = ""
		}

		str += string(data)
	}

	if len(str) != 0 {
		fmt.Printf("read: %s\n", str)
	}
}
