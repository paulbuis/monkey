package ast

import (
	"bytes"
	"monkey/ast"
	"monkey/token"
)

// conforms to interface fmt.Stringer
//
// conforms to interface ast.PrefixExpression
type PrefixExpression struct {
	token          token.Token // The prefix token, e.g. !
	prefixOperator string
	operand        ast.Expression
}

// diagnostic check to verify *PrefixExpression struct
// in this package conforms to ast.PrefixExpression interface
var _ ast.PrefixExpression = &PrefixExpression{}

func NewPrefixExpression(token token.Token, prefixOperator string, operand ast.Expression) *PrefixExpression {
	return &PrefixExpression{token: token, prefixOperator: prefixOperator, operand: operand}
}

func (pe *PrefixExpression) Token() token.Token {
	return pe.token
}

func (pe *PrefixExpression) PrefixOperator() string {
	return pe.prefixOperator
}

func (pe *PrefixExpression) Operand() ast.Expression {
	return pe.operand
}

func (pe *PrefixExpression) ExpressionNode() {}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.token.Literal()
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.prefixOperator)
	out.WriteString(pe.operand.String())
	out.WriteString(")")

	return out.String()
}
