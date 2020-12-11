package ast

import (
	"bytes"
	"monkey/token"
)

// conforms to interface Node
// conforms to interface Expression
// conforms to interface fmt.Stringer
type IndexExpression struct {
	token token.Token // The [ token
	left  Expression
	index Expression
}

func NewIndexExpression(token token.Token,
	left Expression,
	index Expression) *IndexExpression {
	return &IndexExpression{token: token, left: left, index: index}
}

func (ie *IndexExpression) Token() token.Token {
	return ie.token
}

func (ie *IndexExpression) Left() Expression {
	return ie.left
}

func (ie *IndexExpression) Index() Expression {
	return ie.index
}

func (ie *IndexExpression) expressionNode() {}

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
