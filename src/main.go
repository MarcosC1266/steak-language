package main

import (
	"fmt"
	"os"
	"steak/src/file"
	"steak/src/scripts"
	"steak/src/token"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		fmt.Println("No Command-line arguments provided.")
		return
	}

	filePath := args[1]
	dat, err := os.ReadFile(filePath)

	file.CheckFile(err)

	contents := string(dat)

	tokens := token.Tokenize(contents)
	asm := token.TokensToAsm(tokens)
	fileCreation := file.CreateAsmFile(asm)

	if fileCreation {
		fmt.Println("Archivo creado correctamente")
		fmt.Println("Creando archivo de ejecucion")
		scripts.CreateObjectFile()
	}

}
