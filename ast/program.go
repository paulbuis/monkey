package ast

import (
	"bytes"
	"encoding/json"
)

// represents a sequence of ast.Statement
//  conforms to interface Stringer
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
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
