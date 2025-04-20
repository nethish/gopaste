package main

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || !os.IsNotExist(err)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("you didn't provide a filename")
		return
	}
	filename := os.Args[1]

	if fileExists(filename) {
		fmt.Println("file already exists. Enter new filename or delete the file")
		return
	}

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
