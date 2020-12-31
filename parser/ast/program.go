package ast

import (
	"bytes"
	"encoding/json"
	"monkey/ast"
)

//  *Program conforms to interface ast.Program
type Program struct {
	statements []ast.Statement
}

// diagnostic check to verify *LetStatement struct
// in this package conforms to ast.LetStatement interface
var _ ast.Program = &Program{}

func NewProgram(statements []ast.Statement) ast.Program {
	return &Program{statements: statements}
}

func (p *Program) TokenLiteral() string {
	if len(p.statements) > 0 {
		return p.statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) StatementNode() {}

func (p *Program) Statements() []ast.Statement {
	return p.statements
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (p *Program) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["NodeType"] = "Program"
	m["program"] = p.Statements

	return json.Marshal(m)
}
