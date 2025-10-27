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
