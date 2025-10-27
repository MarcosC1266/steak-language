package token

import (
	"fmt"
	"strings"
	"unicode"
)

type TokenType int

const (
	_return TokenType = iota
	int_lit
	semi
)

type Token struct {
	token TokenType
	value string
}

func Tokenize(str string) []Token {
	buf := ""
	strRune := []rune(str)
	var tokens []Token
	for i := 0; i < len(strRune); i++ {
		char := strRune[i]
		var builder strings.Builder
		if unicode.IsLetter(char) {
			builder.WriteRune(char)
			i++
			for unicode.IsLetter(strRune[i]) {
				builder.WriteRune(strRune[i])
				i++
			}
			i--
			buf = builder.String()
			if buf == "return" {
				tokens = append(tokens, Token{_return, buf})
				buf = ""
				continue
			} else {
				fmt.Println("La cagaste compadre")
			}
		} else if unicode.IsDigit(char) {
			builder.WriteRune(char)
			i++
			for unicode.IsDigit(strRune[i]) {
				builder.WriteRune(strRune[i])
				i++
			}
			i--
			buf = builder.String()
			tokens = append(tokens, Token{int_lit, buf})
			buf = ""
			continue
		} else if string(char) == ";" {
			tokens = append(tokens, Token{semi, ""})
		} else if unicode.IsSpace(char) {
			continue
		} else {
			fmt.Println("No paso nada")
		}
	}

	return tokens
}

// todo mejorar esta funcion, ahora esta hard code para tener un output
func TokensToAsm(tokens []Token) string {
	var output strings.Builder
	output.WriteString("global _start\n_start:\n")
	firstToken := tokens[0]
	secondToken := tokens[1]

	if firstToken.token == _return && secondToken.token == int_lit {
		output.WriteString("  mov rax, 60\n")
		output.WriteString("  mov rdi, " + secondToken.value + "\n")
		output.WriteString("  syscall")
	}

	return output.String()

}
