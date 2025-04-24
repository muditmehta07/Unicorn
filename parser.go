package main

import (
	"fmt"
	"strconv"
)

type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens, pos: 0}
}

func (p *Parser) ParseStatement() Stmt {
	tok := p.peek()

	switch tok.Type {

	case TOKEN_PRINT:
		p.next()
		val := p.parseExpression()
		return &PrintStmt{Value: val}

	case TOKEN_IDENTIFIER:
		name := tok.Literal
		p.next()
		p.expect(TOKEN_ASSIGN)
		val := p.parseExpression()
		return &AssignStmt{Name: name, Value: val}

	default:
		panic(fmt.Sprintf("Unknown statement: %v", tok))

	}
}

func (p *Parser) parseExpression() Expr {
	left := p.parsePrimary()

	if p.match(TOKEN_PLUS) {
		right := p.parsePrimary()
		return &BinaryExpr{Left: left, Operator: "+", Right: right}
	}

	return left
}

func (p *Parser) parsePrimary() Expr {
	tok := p.next()

	switch tok.Type {
	case TOKEN_NUMBER:
		val, _ := strconv.Atoi(tok.Literal)
		return &IntegerLiteral{Value: val}
	case TOKEN_IDENTIFIER:
		return &Identifier{Name: tok.Literal}
	default:
		panic(fmt.Sprintf("Invalid primary expression: %v", tok))
	}
}

func (p *Parser) expect(t TokenType) {
	if p.peek().Type == t {
		p.next()
	} else {
		panic(fmt.Sprintf("Expected token %v, got %v", t, p.peek()))
	}
}

func (p *Parser) next() Token {
	if p.pos < len(p.tokens) {
		tok := p.tokens[p.pos]
		p.pos++
		return tok
	}
	return Token{Type: TOKEN_EOF}
}

func (p *Parser) peek() Token {
	if p.pos < len(p.tokens) {
		return p.tokens[p.pos]
	}
	return Token{Type: TOKEN_EOF}
}