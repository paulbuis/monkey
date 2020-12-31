package ast

import (
	"monkey/ast"
	"monkey/token"
)

// *IntegerLiteral conforms to interface ast.IntegerLiteral
//
// conforms to interface fmt.Stringer
type IntegerLiteral struct {
	token token.Token
	value int64
}

// diagnostic check to verify *IntegerLiteral struct
// in this package conforms to ast.IntegerLiteral interface
var _ ast.IntegerLiteral = &IntegerLiteral{}

func NewIntegerLiteral(token token.Token, value int64) *IntegerLiteral {
	return &IntegerLiteral{token: token, value: value}
}

func (il *IntegerLiteral) Token() token.Token {
	return il.token
}

func (il *IntegerLiteral) IntegerValue() int64 {
	return il.value
}

func (il *IntegerLiteral) ExpressionNode() {}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.token.Literal()
}

func (il *IntegerLiteral) String() string {
	return il.token.Literal()
}
