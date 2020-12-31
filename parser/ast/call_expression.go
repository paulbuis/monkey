// Package monkey/parser/ast defines types conforming to interfaces with same names in monkey/ast
package ast

import (
	"bytes"
	"monkey/ast"
	"monkey/token"
	"strings"
)

type CallExpression struct {
	token     token.Token    // The '(' token
	function  ast.Expression // Identifier or FunctionLiteral
	arguments []ast.Expression
}

// diagnostic check to see if *CallExpression struct
// in this package conforms to ast.CallExpression interface
var _ ast.CallExpression = &CallExpression{}

func NewCallExpression(token token.Token, function ast.Expression, arguments []ast.Expression) *CallExpression {
	return &CallExpression{token: token, function: function, arguments: arguments}
}

func (ce *CallExpression) Token() token.Token {
	return ce.token
}

func (ce *CallExpression) Function() ast.Expression {
	return ce.function
}

func (ce *CallExpression) Argument(index int) ast.Expression {
	return ce.arguments[index]
}

func (ce *CallExpression) ArgumentCount() int {
	return len(ce.arguments)
}

func (ce *CallExpression) ExpressionNode() {}

func (ce *CallExpression) TokenLiteral() string {
	return ce.token.Literal()
}

func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := make([]string, len(ce.arguments))
	for i, a := range ce.arguments {
		args[i] = a.String()
	}

	out.WriteString(ce.function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
