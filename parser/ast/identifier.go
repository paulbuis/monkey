package ast

import (
	"monkey/ast"
	"monkey/token"
)

//     conforms to interface Node
//     conforms to interface Expression
//     conforms to interface fmt.Stringer
type Identifier struct {
	token token.Token // the token.IDENT token
	value string
}

// diagnostic check to see if *HashLiteral struct
// in this package conforms to ast.HashLiteral interface
var _ ast.Identifier = &Identifier{}

func NewIdentifier(token token.Token, value string) *Identifier {
	return &Identifier{token: token, value: value}
}

func (i *Identifier) Token() token.Token {
	return i.token
}

func (i *Identifier) IdentifierName() string {
	return i.value
}

func (i *Identifier) ExpressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.token.Literal()
}

func (i *Identifier) String() string {
	return i.value
}
