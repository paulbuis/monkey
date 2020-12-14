package ast

import (
	"bytes"
	"encoding/json"
	"monkey/token"
)

// conforms to interface Node
// conforms to interface Expression
// conforms to interface fmt.Stringer
type InfixExpression struct {
	token    token.Token // The operator token, e.g. +
	left     Expression
	operator string
	right    Expression
}

func NewInfixExpression(token token.Token, left Expression, operator string, right Expression) *InfixExpression {
	return &InfixExpression{token: token, left: left, operator: operator, right: right}
}

func (ie *InfixExpression) Token() token.Token {
	return ie.token
}

func (ie *InfixExpression) Left() Expression {
	return ie.left
}

func (ie *InfixExpression) Operator() string {
	return ie.operator
}

func (ie *InfixExpression) Right() Expression {
	return ie.right
}

func (ie *InfixExpression) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["NodeType"] = "InfixExpression"
	m[" token"] = ie.token
	m["left"] = ie.left
	m["operator"] = ie.operator
	m["right"] = ie.right

	return json.Marshal(m)
}

func (ie *InfixExpression) expressionNode() {}

func (ie *InfixExpression) TokenLiteral() string {
	return ie.token.Literal()
}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.left.String())
	out.WriteString(" " + ie.operator + " ")
	out.WriteString(ie.right.String())
	out.WriteString(")")

	return out.String()
}
