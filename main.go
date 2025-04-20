package main

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("you didn't provide a filename")
		return
	}
	filename := os.Args[1]
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("error when opening file:", err)
	}
	defer f.Close()

	text, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("failed to read clipboard:", err)
		return
	}

	f.Write([]byte(text))
}
