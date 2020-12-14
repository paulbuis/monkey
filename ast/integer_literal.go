package ast

import (
	"encoding/json"
	"monkey/token"
)

// conforms to interface Node
// conforms to interface Expression
// conforms to interface fmt.Stringer
type IntegerLiteral struct {
	token token.Token
	value int64
}

func NewIntegerLiteral(token token.Token, value int64) *IntegerLiteral {
	return &IntegerLiteral{token: token, value: value}
}

func (il *IntegerLiteral) Token() token.Token {
	return il.token
}

func (il *IntegerLiteral) Value() int64 {
	return il.value
}

func (il *IntegerLiteral) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["NodeType"] = "IntegerLiteral"
	m[" token"] = il.token

	return json.Marshal(m)
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.token.Literal()
}

func (il *IntegerLiteral) String() string {
	return il.token.Literal()
}
