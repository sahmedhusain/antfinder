package main

import (
	f "antfinder/functions"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		f.ProcessFile(os.Args[1])
	} else {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}
}
