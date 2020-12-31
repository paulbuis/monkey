package ast

import (
	"bytes"
	"monkey/ast"
	"monkey/token"
)

// conforms to interface fmt.Stringer
// *IfExpression conforms to interface ast.IfExpression
type IfExpression struct {
	// public read-only via IfExpression.Token() method
	token token.Token // The 'if' token
	// public read-only via IfExpression.Condition() method
	// always non-nil
	condition ast.Expression
	// public read-only via IfExpression.Consequence() method
	// always non-nil
	consequence ast.BlockStatement
	// public read-Only via IfExpression.Alternative() method
	// may be nil
	alternative ast.BlockStatement
}

// diagnostic check to see if *IfExpression struct
// in this package conforms to ast.IfExpression interface
var _ ast.IfExpression = &IfExpression{}

func NewIfExpression(token token.Token,
	condition ast.Expression,
	consequence ast.BlockStatement,
	alternative ast.BlockStatement) *IfExpression {
	return &IfExpression{token: token,
		condition:   condition,
		consequence: consequence,
		alternative: alternative}
}

func (ie *IfExpression) Token() token.Token {
	return ie.token
}

func (ie *IfExpression) Condition() ast.Expression {
	return ie.condition
}

func (ie *IfExpression) Consequence() ast.BlockStatement {
	return ie.consequence
}

func (ie *IfExpression) Alternative() ast.BlockStatement {
	return ie.alternative
}

func (ie *IfExpression) ExpressionNode() {}

func (ie *IfExpression) TokenLiteral() string {
	return ie.token.Literal()
}

func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if (")
	out.WriteString(ie.condition.String())
	out.WriteString(") ")
	out.WriteString(ie.consequence.String())

	if ie.alternative != nil {
		out.WriteString(" else ")
		out.WriteString(ie.alternative.String())
	}

	return out.String()
}
