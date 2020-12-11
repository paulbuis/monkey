package ast

import (
	"monkey/token"
)

// represents an expression statement
//    conforms to interface Node
//    conforms to interface Statement
//	  conforms to interface fmt.Stringer
type ExpressionStatement struct {
	token      token.Token // the first token of the expression
	expression Expression
}

func NewExpressionStatement(token token.Token, expression Expression) *ExpressionStatement {
	return &ExpressionStatement{token: token, expression: expression}
}

func (es *ExpressionStatement) Token() token.Token {
	return es.token
}

func (es *ExpressionStatement) Expression() Expression {
	return es.expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.token.Literal()
}

func (es *ExpressionStatement) String() string {
	if es.Expression() != nil {
		return es.expression.String()
	}
	return ""
}
