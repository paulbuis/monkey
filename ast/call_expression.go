package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

type CallExpression struct {
	token     token.Token // The '(' token
	function  Expression  // Identifier or FunctionLiteral
	arguments []Expression
}

func NewCallExpression(token token.Token, function Expression, arguments []Expression) *CallExpression {
	return &CallExpression{token: token, function: function, arguments: arguments}
}

func (ce *CallExpression) Token() token.Token {
	return ce.token
}

func (ce *CallExpression) Function() Expression {
	return ce.function
}

func (ce *CallExpression) Arguments() []Expression {
	return ce.arguments
}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) TokenLiteral() string {
	return ce.token.Literal()
}

func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := make([]string, 0, len(ce.arguments))
	for _, a := range ce.arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
