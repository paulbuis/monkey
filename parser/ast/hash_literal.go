package ast

import (
	"bytes"
	"monkey/ast"
	"monkey/token"
	"strings"
)

// conforms to interface Expression
// conforms to interface Node
// conforms to interface fmt.Stringer
type HashLiteral struct {
	token  token.Token // the '{' token
	keys   []ast.Expression
	values []ast.Expression
}

// diagnostic check to see if *HashLiteral struct
// in this package conforms to ast.HashLiteral interface
var _ ast.HashLiteral = &HashLiteral{}

func NewHashLiteral(token token.Token, keys []ast.Expression, values []ast.Expression) *HashLiteral {
	return &HashLiteral{token: token, keys: keys, values: values}
}

func (hl *HashLiteral) Token() token.Token {
	return hl.token
}

func (hl *HashLiteral) PairCount() int {
	return len(hl.keys)
}

func (hl *HashLiteral) Keys(index int) ast.Expression {
	return hl.keys[index]
}

func (hl *HashLiteral) Values(index int) ast.Expression {
	return hl.values[index]
}

func (hl *HashLiteral) ExpressionNode() {}

func (hl *HashLiteral) TokenLiteral() string {
	return hl.token.Literal()
}

func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := make([]string, len(hl.keys))
	for i, key := range hl.keys {
		value := hl.values[i]
		pairs[i] = key.String() + ":" + value.String()
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	result := out.String()
	return result
}
