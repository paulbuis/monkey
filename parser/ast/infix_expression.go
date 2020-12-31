package ast

import (
	"bytes"
	"monkey/ast"
	"monkey/token"
)

// conforms to interface fmt.Stringer
// *InfixExpression conforms to interface ast.InfixExpression
type InfixExpression struct {
	token    token.Token // The operator token, e.g. +
	left     ast.Expression
	operator string
	right    ast.Expression
}

// diagnostic check to verify *InfixExpression struct
// in this package conforms to ast.InfixExpression interface
var _ ast.InfixExpression = &InfixExpression{}
var _ ast.Expression = &InfixExpression{}
var _ ast.Node = &InfixExpression{}

func NewInfixExpression(token token.Token,
	left ast.Expression,
	operator string,
	right ast.Expression,
) *InfixExpression {
	return &InfixExpression{token: token, left: left, operator: operator, right: right}
}

func (ie *InfixExpression) Token() token.Token {
	return ie.token
}

func (ie *InfixExpression) Left() ast.Expression {
	return ie.left
}

func (ie *InfixExpression) Operator() string {
	return ie.operator
}

func (ie *InfixExpression) Right() ast.Expression {
	return ie.right
}

func (ie *InfixExpression) ExpressionNode() {}

func (ie *InfixExpression) TokenLiteral() string {
	return ie.token.Literal()
}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.left.String())
	out.WriteString(" " + ie.operator + " ")
	out.WriteString(ie.right.String())
	out.WriteString(")")

	return out.String()
}
