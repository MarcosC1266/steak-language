package token

import "fmt"

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

func Tokenize(str string) Token {

	for _, r := range str {
		char := string(r)
		fmt.Println(char)
	}

	result := Token{1, "hola"}
	return result
}
