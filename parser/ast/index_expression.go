package ast

import (
	"bytes"
	"monkey/ast"
	"monkey/token"
)

// conforms to interface fmt.Stringer
//
// conforms to interface ast.IndexExpression
type IndexExpression struct {
	// public read-only via IndexExpression.Token() method
	token token.Token // The [ token
	// public read-only via IndexExpression.Left() method
	left ast.Expression
	// public read-only via IndexExpression.Index() method
	index ast.Expression
}

// diagnostic check to see if *IndexExpression struct
// in this package conforms to ast.IndexExpression interface
var _ ast.IndexExpression = &IndexExpression{}

func NewIndexExpression(token token.Token,
	left ast.Expression,
	index ast.Expression,
) *IndexExpression {
	return &IndexExpression{token: token, left: left, index: index}
}

func (ie *IndexExpression) Token() token.Token {
	return ie.token
}

func (ie *IndexExpression) Left() ast.Expression {
	return ie.left
}

func (ie *IndexExpression) Index() ast.Expression {
	return ie.index
}

func (ie *IndexExpression) ExpressionNode() {}

func (ie *IndexExpression) TokenLiteral() string {
	return ie.token.Literal()
}

func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.left.String())
	out.WriteString("[")
	out.WriteString(ie.index.String())
	out.WriteString("])")

	return out.String()
}
