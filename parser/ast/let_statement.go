package ast

import (
	"bytes"
	"monkey/ast"
	"monkey/token"
)

//  *LetStatement struct conforms to interface ast.LetStatement
type LetStatement struct {
	token token.Token // the token.LET token
	name  *Identifier
	value ast.Expression
}

// diagnostic check to verify *LetStatement struct
// in this package conforms to ast.LetStatement interface
var _ ast.LetStatement = &LetStatement{}

//
func NewLetStatement(
	token token.Token,
	name *Identifier,
	value ast.Expression,
) *LetStatement {
	return &LetStatement{token: token, name: name, value: value}
}

func (ls *LetStatement) Token() token.Token {
	return ls.token
}

func (ls *LetStatement) Name() ast.Identifier {
	return ls.name
}

func (ls *LetStatement) LetValue() ast.Expression {
	return ls.value
}

func (ls *LetStatement) StatementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.token.Literal()
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.name.String())
	out.WriteString(" = ")

	if ls.value != nil {
		out.WriteString(ls.value.String())
	}

	out.WriteString(";")

	return out.String()
}
