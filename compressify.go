package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Printf("Usage instructions:\n%v <file-to-compress>\n", args[0])
		return
	}

	file, err := os.Open(args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	bytes := make([]byte, 100)

	n, err := file.Read(bytes)
	if err != nil {
		fmt.Println("Error reading bytes:", err)
		return
	}
	
	c := 1
	var sb strings.Builder

	for i, byte := range bytes {
		if i + 1 < len(bytes) && byte == bytes[i + 1] {
			c++
		} else {
			sb.WriteByte(byte)
		}
	}

	result := sb.String()
	fmt.Println(result)

	fmt.Printf("Read %d bytes: %v\n", n, bytes)
}
