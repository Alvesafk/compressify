package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Printf("Usage instructions:\n%v <file-to-compress>\n", args[0])
		return
	}

	fmt.Println("Hello, world!")
}
