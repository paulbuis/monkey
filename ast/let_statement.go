package ast

import (
	"bytes"
	"encoding/json"
	"monkey/token"
)

// represents a let statement
//  conforms to interface Statement
//  conforms to interface Node
//  conforms to interface fmt.Stringer
type LetStatement struct {
	token token.Token // the token.LET token
	name  *Identifier
	value Expression
}

func NewLetStatement(
	token token.Token,
	name *Identifier,
	value Expression,
) *LetStatement {
	return &LetStatement{token: token, name: name, value: value}
}

func (ls *LetStatement) Token() token.Token {
	return ls.token
}

func (ls *LetStatement) Name() *Identifier {
	return ls.name
}

func (ls *LetStatement) Value() Expression {
	return ls.value
}

func (ls *LetStatement) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["NodeType"] = "LetStatement"
	m[" token"] = ls.token
	m["name"] = ls.name
	m["value"] = ls.value

	return json.Marshal(m)
}
func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.token.Literal()
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.name.String())
	out.WriteString(" = ")

	if ls.value != nil {
		out.WriteString(ls.value.String())
	}

	out.WriteString(";")

	return out.String()
}
