package ast

import (
	"bytes"
	"monkey/ast"
	"monkey/token"
)

// *ReturnStatement conforms to interface ast.ReturnStatement
type ReturnStatement struct {
	token       token.Token // the 'return' token
	returnValue ast.Expression
}

// diagnostic check to verify *ReturnStatement struct
// in this package conforms to ast.ReturnStatement interface
var _ ast.ReturnStatement = &ReturnStatement{}

// this generates a compile error, which is a good thing!!!
// var _ ast.ReturnStatement = &ExpressionStatement{}

func NewReturnStatement(token token.Token, returnValue ast.Expression) *ReturnStatement {
	return &ReturnStatement{token: token, returnValue: returnValue}
}

func (rs *ReturnStatement) Token() token.Token {
	return rs.token
}

func (rs *ReturnStatement) ReturnValue() ast.Expression {
	return rs.returnValue
}

func (rs *ReturnStatement) StatementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.token.Literal() }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.returnValue != nil {
		out.WriteString(rs.returnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
