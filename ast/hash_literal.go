package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

// conforms to interface Expression
// conforms to interface Node
// conforms to interface fmt.Stringer
type HashLiteral struct {
	token token.Token // the '{' token
	pairs map[Expression]Expression
}

func NewHashLiteral(token token.Token, pairs map[Expression]Expression) *HashLiteral {
	return &HashLiteral{token: token, pairs: pairs}
}

func (hl *HashLiteral) Pairs() map[Expression]Expression {
	return hl.pairs
}

func (hl *HashLiteral) Token() token.Token {
	return hl.token
}

func (hl *HashLiteral) expressionNode() {}

func (hl *HashLiteral) TokenLiteral() string {
	return hl.token.Literal()
}

func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := make([]string, 0, len(hl.pairs))
	for key, value := range hl.pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
