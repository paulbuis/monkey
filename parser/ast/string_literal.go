package ast

import (
	"monkey/ast"
	"monkey/token"
)

// *StringLiteral conforms to interface ast.StringLiteral
type StringLiteral struct {
	token token.Token
	value string
}

// diagnostic check to verify *StringLiteral struct
// in this package conforms to ast.StringLiteral interface
var _ ast.StringLiteral = &StringLiteral{}

func NewStringLiteral(token token.Token, value string) *StringLiteral {
	return &StringLiteral{token: token, value: value}
}

func (sl *StringLiteral) Token() token.Token {
	return sl.token
}

func (sl *StringLiteral) StringValue() string {
	return sl.value
}

func (sl *StringLiteral) ExpressionNode() {}

func (sl *StringLiteral) TokenLiteral() string {
	return sl.token.Literal()
}

func (sl *StringLiteral) String() string {
	return sl.token.Literal()
}
