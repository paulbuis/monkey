package ast

import (
	"bytes"
	"encoding/json"
	"monkey/token"
)

type ReturnStatement struct {
	token       token.Token // the 'return' token
	returnValue Expression
}

func NewReturnStatement(token token.Token, returnValue Expression) *ReturnStatement {
	return &ReturnStatement{token: token, returnValue: returnValue}
}

func (rs *ReturnStatement) Token() token.Token {
	return rs.token
}

func (rs *ReturnStatement) ReturnValue() Expression {
	return rs.returnValue
}

func (rs *ReturnStatement) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["NodeType"] = "ReturnStatement"
	m[" token"] = rs.token
	m["returnValue"] = rs.returnValue

	return json.Marshal(m)
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.token.Literal() }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.returnValue != nil {
		out.WriteString(rs.returnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
