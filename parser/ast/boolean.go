package ast

import (
	"monkey/ast"
	"monkey/token"
)

// conforms to interface Node
// conforms to interface Expression
// conforms to interface fmt.Stringer
type Boolean struct {
	token token.Token
	value bool
}

// diagnostic check to see if *Boolean struct
// in this package conforms to ast.Boolean interface
var _ ast.Boolean = &Boolean{}

func NewBoolean(token token.Token, value bool) *Boolean {
	return &Boolean{token: token, value: value}
}

func (b *Boolean) Token() token.Token {
	return b.token
}

func (b *Boolean) BooleanValue() bool {
	return b.value
}

func (b *Boolean) ExpressionNode() {}

func (b *Boolean) TokenLiteral() string {
	return b.token.Literal()
}

func (b *Boolean) String() string {
	return b.token.Literal()
}
