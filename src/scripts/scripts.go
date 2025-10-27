package scripts

import (
	"os/exec"
	"steak/src/file"
)

func CreateObjectFile() {
	fileObject := exec.Command("nasm", "-felf64", "./result/out.asm")
	execFile := exec.Command("ld", "-o", "./result/out", "./result/out.o")

	_, err := fileObject.Output()
	file.CheckFile(err)

	_, err = execFile.Output()
	file.CheckFile(err)

}
