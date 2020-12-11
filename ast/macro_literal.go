package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

type MacroLiteral struct {
	token      token.Token // The 'macro' token
	parameters []*Identifier
	body       *BlockStatement
}

func NewMacroLiteral(
	token token.Token,
	parameters []*Identifier,
	body *BlockStatement,
) *MacroLiteral {
	return &MacroLiteral{token: token, parameters: parameters, body: body}
}

func (ml *MacroLiteral) Token() token.Token {
	return ml.token
}

func (ml *MacroLiteral) Parameters() []*Identifier {
	return ml.parameters
}

func (ml *MacroLiteral) Body() *BlockStatement {
	return ml.body
}

func (ml *MacroLiteral) expressionNode()      {}
func (ml *MacroLiteral) TokenLiteral() string { return ml.token.Literal() }
func (ml *MacroLiteral) String() string {
	var out bytes.Buffer

	params := make([]string, 0, len(ml.parameters))
	for _, p := range ml.parameters {
		params = append(params, p.String())
	}

	out.WriteString(ml.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(ml.body.String())

	return out.String()
}
