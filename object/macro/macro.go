package macro

import (
	"bytes"
	"monkey/ast"
	"monkey/object"
	"strings"
)

type Macro struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *object.Environment
}

func (m *Macro) Type() object.ObjectType {
	return object.MACRO_OBJ
}

func (m *Macro) Inspect() string {
	var out bytes.Buffer

	params := make([]string, 0, len(m.Parameters))
	for _, p := range m.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("macro")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(m.Body.String())
	out.WriteString("\n}")

	return out.String()
}
