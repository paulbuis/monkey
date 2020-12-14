package ast

import (
	"bytes"
	"encoding/json"
	"monkey/token"
)

// conforms to interface Node
// conforms to interface Expression
// conforms to interface fmt.Stringer
type PrefixExpression struct {
	token    token.Token // The prefix token, e.g. !
	operator string
	right    Expression
}

func NewPrefixExpression(token token.Token, operator string, right Expression) *PrefixExpression {
	return &PrefixExpression{token: token, operator: operator, right: right}
}

func (pe *PrefixExpression) Token() token.Token {
	return pe.token
}

func (pe *PrefixExpression) Operator() string {
	return pe.operator
}

func (pe *PrefixExpression) Right() Expression {
	return pe.right
}

func (pe *PrefixExpression) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["NodeType"] = "PrefixExpression"
	m[" token"] = pe.token
	m["operator"] = pe.operator
	m["right"] = pe.right
	return json.Marshal(m)
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.token.Literal()
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.operator)
	out.WriteString(pe.right.String())
	out.WriteString(")")

	return out.String()
}
