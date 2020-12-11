package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

type ArrayLiteral struct {
	token    token.Token // the '[' token
	elements []Expression
}

func NewArrayLiteral(token token.Token, elements []Expression) *ArrayLiteral {
	return &ArrayLiteral{token: token, elements: elements}
}

func (al *ArrayLiteral) Token() token.Token {
	return al.token
}

func (al *ArrayLiteral) Elements() []Expression {
	return al.elements
}

func (al *ArrayLiteral) expressionNode() {}

func (al *ArrayLiteral) TokenLiteral() string {
	return al.token.Literal()
}

func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := make([]string, 0, len(al.elements))
	for _, el := range al.elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
