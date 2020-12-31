package ast

import (
	"bytes"
	"monkey/ast"
	"monkey/token"
	"strings"
)

type FunctionLiteral struct {
	token      token.Token // The 'fn' token
	parameters []ast.Identifier
	body       ast.BlockStatement
}

// diagnostic check to see if *FunctionLiteral struct
// in this package conforms to ast.FunctionLiteral interface
var _ ast.FunctionLiteral = &FunctionLiteral{}

func NewFunctionLiteral(token token.Token, parameters []ast.Identifier, body ast.BlockStatement) ast.FunctionLiteral {
	return &FunctionLiteral{token: token, parameters: parameters, body: body}
}

func (fl *FunctionLiteral) Token() token.Token {
	return fl.token
}

func (fl *FunctionLiteral) Parameter(index int) ast.Identifier {
	return fl.parameters[index]
}

func (fl *FunctionLiteral) ParameterCount() int {
	return len(fl.parameters)
}

func (fl *FunctionLiteral) Body() ast.BlockStatement {
	return fl.body
}

func (fl *FunctionLiteral) ExpressionNode() {}

func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.token.Literal()
}

func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := make([]string, len(fl.parameters))
	for i, p := range fl.parameters {
		params[i] = p.String()
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.body.String())

	return out.String()
}
