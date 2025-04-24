package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var filename string
	if len(os.Args) != 2 {
		fmt.Println("Usage: unicorn <filename>.uni")
		os.Exit(1)
	}
	filename = os.Args[1]

	if len(filename) < 4 || filename[len(filename)-4:] != ".uni" {
		fmt.Println("Error: Please provide a .uni file.")
		os.Exit(1)
	}

	code, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		os.Exit(1)
	}

	lines := splitLines(string(code))
	env := NewEnvironment()

	for _, line := range lines {
		tokens := Lex(line)
		parser := NewParser(tokens)
		stmt := parser.ParseStatement()
		env.Eval(stmt)
	}
}

func splitLines(code string) []string {
	var result []string
	current := ""
	for _, ch := range code {
		if ch == '\n' {
			result = append(result, current)
			current = ""
		} else {
			current += string(ch)
		}
	}

	if current != "" {
		result = append(result, current)
	}

	return result
}
