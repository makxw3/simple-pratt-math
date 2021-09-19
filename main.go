package main

import (
	"fmt"
	"pratt/lexer"
	"pratt/parser"
)

func main() {
	inp := "1 + 2 + 3 + 4 * 5 - 6"
	lx := lexer.New(inp)
	ps := parser.New(lx)
	exp := ps.ParseMathExpression()

	fmt.Println(exp.ExpressionString())
}
