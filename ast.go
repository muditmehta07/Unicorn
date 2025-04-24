package main

import "fmt"

type Node interface {
	String() string
}

type Expr interface {
	Node
	exprNode()
}

type Stmt interface {
	Node
	stmtNode()
}

type IntegerLiteral struct {
	Value int
}

func (i *IntegerLiteral) exprNode() {}

func (i *IntegerLiteral) String() string {
	return fmt.Sprintf("%d", i.Value)
}

type Identifier struct {
	Name string
}

func (i *Identifier) exprNode() {}

func (i *Identifier) String() string {
	return i.Name
}

type BinaryExpr struct {
	Left     Expr
	Operator string
	Right    Expr
}

func (b *BinaryExpr) exprNode() {}

func (b *BinaryExpr) String() string {
	return fmt.Sprintf("(%s %s %s)", b.Left.String(), b.Operator, b.Right.String())
}

type PrintStmt struct {
	Value Expr
}

func (p *PrintStmt) stmtNode() {}

func (p *PrintStmt) String() string {
	return "print " + p.Value.String()
}

type AssignStmt struct {
	Name  string
	Value Expr
}

func (a *AssignStmt) stmtNode() {}

func (a *AssignStmt) String() string {
	return fmt.Sprintf("%s = %s", a.Name, a.Value.String())
}