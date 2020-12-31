package ast

import (
	"bytes"
	"monkey/ast"
	"monkey/token"
	"strings"
)

// diagnostic check to see if *ArrayLiteral struct
// in this package conforms to ast.ArrayLiteral interface
var _ ast.ArrayLiteral = &ArrayLiteral{}

// *ArrayLiteral conforms to interface ast.ArrayLiteral
type ArrayLiteral struct {
	token    token.Token // the '[' token
	elements []ast.Expression
}

func NewArrayLiteral(token token.Token, elements []ast.Expression) *ArrayLiteral {
	return &ArrayLiteral{token: token, elements: elements}
}

func (al *ArrayLiteral) Token() token.Token {
	return al.token
}

func (al *ArrayLiteral) Elements() []ast.Expression {
	return al.elements
}

func (al *ArrayLiteral) ExpressionNode() {}

func (al *ArrayLiteral) TokenLiteral() string {
	return al.token.Literal()
}

func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := make([]string, len(al.elements))
	for i, el := range al.elements {
		elements[i] = el.String()
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
