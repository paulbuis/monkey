package ast

import (
	"encoding/json"
	"monkey/token"
)

// conforms to interface Node
// conforms to interface Expression
// conforms to interface fmt.Stringer
type StringLiteral struct {
	token token.Token
	value string
}

func NewStringLiteral(token token.Token, value string) *StringLiteral {
	return &StringLiteral{token: token, value: value}
}

func (sl *StringLiteral) Token() token.Token {
	return sl.token
}

func (sl *StringLiteral) Value() string {
	return sl.value
}

func (sl *StringLiteral) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["NodeType"] = "StringLiteral"
	m[" token"] = sl.token
	m["value"] = sl.value

	return json.Marshal(m)
}

func (sl *StringLiteral) expressionNode() {}

func (sl *StringLiteral) TokenLiteral() string {
	return sl.token.Literal()
}

func (sl *StringLiteral) String() string {
	return sl.token.Literal()
}
