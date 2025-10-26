package main

import (
	"fmt"
	"os"
	"steak/src/token"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		fmt.Println("No Command-line arguments provided.")
		return
	}

	filePath := args[1]
	file, err := os.ReadFile(filePath)

	checkError(err)

	contents := string(file)

	token.Tokenize(contents)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
