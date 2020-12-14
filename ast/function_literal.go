package ast

import (
	"bytes"
	"encoding/json"
	"monkey/token"
	"strings"
)

type FunctionLiteral struct {
	token      token.Token // The 'fn' token
	parameters []*Identifier
	body       *BlockStatement
}

func NewFunctionLiteral(token token.Token, parameters []*Identifier, body *BlockStatement) *FunctionLiteral {
	return &FunctionLiteral{token: token, parameters: parameters, body: body}
}

func (fl *FunctionLiteral) Token() token.Token {
	return fl.token
}

func (fl *FunctionLiteral) Parameters() []*Identifier {
	return fl.parameters
}

func (fl *FunctionLiteral) Body() *BlockStatement {
	return fl.body
}

func (fl *FunctionLiteral) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["NodeType"] = "FunctionLiteral"
	m[" token"] = fl.token
	m["arguments"] = fl.parameters
	m["body"] = fl.body

	return json.Marshal(m)
}

func (fl *FunctionLiteral) expressionNode() {}

func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.token.Literal()
}

func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := make([]string, 0, len(fl.parameters))
	for _, p := range fl.parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.body.String())

	return out.String()
}
