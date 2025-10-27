package file

import "os"

func CreateAsmFile(str string) bool {
	fileName := "./result/out.asm"
	file, err := os.Create(fileName)
	CheckFile(err)
	defer file.Close()

	_, err = file.WriteString(str)
	CheckFile(err)
	return true
}
