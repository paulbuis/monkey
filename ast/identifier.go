package ast

import (
	"encoding/json"
	"monkey/token"
)

//     conforms to interface Node
//     conforms to interface Expression
//     conforms to interface fmt.Stringer
type Identifier struct {
	token token.Token // the token.IDENT token
	value string
}

func NewIdentifier(token token.Token, value string) *Identifier {
	return &Identifier{token: token, value: value}
}

func (i *Identifier) Token() token.Token {
	return i.token
}

func (i *Identifier) Value() string {
	return i.value
}

func (i *Identifier) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["NodeType"] = "Identifier"
	m[" token"] = i.token
	m["value"] = i.value
	return json.Marshal(m)
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.token.Literal()
}

func (i *Identifier) String() string {
	return i.value
}
