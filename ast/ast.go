package ast

import (
	"bytes"
	"fmt"
	"pratt/token"
)

// An expression
type Expression interface {
	ExpressionString() string
}

type SingleNumberExpression struct {
	Value int64 // The integer value for the expression
}

func (sne *SingleNumberExpression) ExpressionString() string {
	var out bytes.Buffer
	// out.WriteString("(")
	out.WriteString(fmt.Sprint(sne.Value))
	// out.WriteString(")")
	return out.String()
}

type InfixExpression struct {
	Operator token.Token
	Left     Expression // The left expression
	Right    Expression // The right expression
}

func (ie *InfixExpression) ExpressionString() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.ExpressionString())
	out.WriteString(" " + ie.Operator.Literal + " ")
	out.WriteString(ie.Right.ExpressionString())
	out.WriteString(")")
	return out.String()
}
