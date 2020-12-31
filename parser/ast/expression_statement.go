package ast

import (
	"monkey/ast"
	"monkey/token"
)

// represents an expression statement
//    conforms to interface Node
//    conforms to interface Statement
//	  conforms to interface fmt.Stringer
type ExpressionStatement struct {
	token      token.Token // the first token of the expression
	expression ast.Expression
}

// diagnostic check to see if *ExpressionStatement struct
// in this package conforms to ast.ExpressionStatement  interface
var _ ast.ExpressionStatement = &ExpressionStatement{}

// this generates a compile error, which seems like a good thing!
//var _ ast.ReturnStatement = &ExpressionStatement{}

func NewExpressionStatement(token token.Token, expression ast.Expression) *ExpressionStatement {
	return &ExpressionStatement{token: token, expression: expression}
}

func (es *ExpressionStatement) Token() token.Token {
	return es.token
}

func (es *ExpressionStatement) Expression() ast.Expression {
	return es.expression
}

func (es *ExpressionStatement) StatementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.token.Literal()
}

func (es *ExpressionStatement) String() string {
	if es.Expression() != nil {
		return es.expression.String()
	}
	return ""
}
