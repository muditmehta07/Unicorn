package main

import (
	"strings"
	"unicode"
)

func Lex(input string) []Token {
	var tokens []Token
	i := 0

	for i < len(input) {
		ch := input[i]

		switch {
		case unicode.IsSpace(rune(ch)):
			i++

		case isLetter(ch):
			start := i
			for i < len(input) && isLetterOrDigit(input[i]) {
				i++
			}
			word := input[start:i]
			tokenType := TOKEN_IDENTIFIER
			if strings.ToUpper(word) == "PRINT" {
				tokenType = TOKEN_PRINT
			}
			tokens = append(tokens, Token{Type: tokenType, Literal: word})

		case unicode.IsDigit(rune(ch)):
			start := i
			for i < len(input) && unicode.IsDigit(rune(input[i])) {
				i++
			}
			tokens = append(tokens, Token{Type: TOKEN_NUMBER, Literal: input[start:i]})

		case ch == '=':
			tokens = append(tokens, Token{Type: TOKEN_ASSIGN, Literal: string(ch)})
			i++

		case ch == '+':
			tokens = append(tokens, Token{Type: TOKEN_PLUS, Literal: string(ch)})
			i++
		
		default:
			tokens = append(tokens, Token{Type: TOKEN_ILLEGAL, Literal: string(ch)})
		}
	}

	tokens = append(tokens, Token{Type: TOKEN_EOF, Literal: ""})
	return tokens
}

func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch))
}

func isLetterOrDigit(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || unicode.IsDigit(rune(ch))
}