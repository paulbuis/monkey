package ast

import (
	"bytes"
	"encoding/json"
	"monkey/token"
	"strings"
)

// conforms to interface Expression
// conforms to interface Node
// conforms to interface fmt.Stringer
type HashLiteral struct {
	token token.Token // the '{' token
	//pairs map[Expression]Expression
	keys   []Expression
	values []Expression
}

func NewHashLiteral(token token.Token, keys []Expression, values []Expression) *HashLiteral {
	return &HashLiteral{token: token, keys: keys, values: values}
}

/*
func (hl *HashLiteral) Pairs() map[Expression]Expression {
	return hl.pairs
}
*/

func (hl *HashLiteral) Token() token.Token {
	return hl.token
}

func (hl *HashLiteral) Keys() []Expression {
	return hl.keys
}

func (hl *HashLiteral) Values() []Expression {
	return hl.values
}

func (hl *HashLiteral) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["NodeType"] = "HashLiteral"
	m[" token"] = hl.token
	m["keys"] = hl.keys
	m["values"] = hl.values

	return json.Marshal(m)
}

func (hl *HashLiteral) expressionNode() {}

func (hl *HashLiteral) TokenLiteral() string {
	return hl.token.Literal()
}

func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := make([]string, 0, len(hl.keys))
	for i, key := range hl.keys {
		value := hl.values[i]
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
