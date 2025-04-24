package main

type TokenType string

const (
	TOKEN_ILLEGAL		TokenType = "ILLEGAL"
	TOKEN_EOF			TokenType = "EOF"
	TOKEN_IDENTIFIER	TokenType = "IDENTIFIER"
	TOKEN_NUMBER		TokenType = "NUMBER"
	TOKEN_ASSIGN		TokenType = "="
	TOKEN_PLUS			TokenType = "+"
	TOKEN_PRINT			TokenType = "PRINT"
)

type Token struct {
	Type 	TokenType
	Literal string
}

func (t Token) String() string {
	return "Token{" + "Type: " + string(t.Type) + ", Literal: " + t.Literal + "}"
}