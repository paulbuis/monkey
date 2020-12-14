package ast

import (
	"encoding/json"
	"monkey/token"
)

// conforms to interface Node
// conforms to interface Expression
// conforms to interface fmt.Stringer
type Boolean struct {
	token token.Token
	value bool
}

func NewBoolean(token token.Token, value bool) *Boolean {
	return &Boolean{token: token, value: value}
}

func (b *Boolean) Token() token.Token {
	return b.token
}

func (b *Boolean) Value() bool {
	return b.value
}

func (b *Boolean) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["NodeType"] = "Boolean"
	m[" token"] = b.token
	m["value"] = b.value

	return json.Marshal(m)
}

func (b *Boolean) expressionNode() {}

func (b *Boolean) TokenLiteral() string {
	return b.token.Literal()
}

func (b *Boolean) String() string {
	return b.token.Literal()
}
