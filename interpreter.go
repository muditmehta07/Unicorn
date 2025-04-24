package main

import "fmt"

type Environment struct {
	variables map[string]int
}

func NewEnvironment() *Environment {
	return &Environment{variables: make(map[string]int)}
}

func (env *Environment) Eval(stmt Stmt) {
	switch stmt := stmt.(type) {
	case *PrintStmt:
		val := env.EvalExpr(stmt.Value)
		fmt.Println(val)
	case *AssignStmt:
		val := env.EvalExpr(stmt.Value)
		env.variables[stmt.Name] = val
	default:
		panic(fmt.Sprintf("Unknown statement: %v", stmt))
	}
}

func (env *Environment) EvalExpr(expr Expr) int {
	switch expr := expr.(type) {
	case *IntegerLiteral:
		return expr.Value
	case *Identifier:
		val, ok := env.variables[expr.Name]
		if !ok {
			panic(fmt.Sprintf("Undefined variable: %s", expr.Name))
		}
		return val
	case *BinaryExpr:
		left := env.EvalExpr(expr.Left)
		right := env.EvalExpr(expr.Right)

		switch expr.Operator {
		case "+":
			return left + right
		default:
			panic(fmt.Sprintf("Unknown operator: %s", expr.Operator))
		}
	default:
		panic(fmt.Sprintf("Unknown expression: %v", expr))
	}
}
