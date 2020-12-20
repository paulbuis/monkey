package macro

import (
	"bytes"
	"monkey/ast"
	"monkey/object"
	"monkey/object/function/environment"
	"strings"
)

type Macro struct {
	Parameters []ast.Identifier
	Body       ast.BlockStatement
	Env        *environment.Environment
}

var _ object.Object = &Macro{}

func (m *Macro) Type() object.ObjectType {
	return object.MACRO_OBJ
}

func (m *Macro) Inspect() string {
	var out bytes.Buffer

	params := make([]string, len(m.Parameters))
	for i, p := range m.Parameters {
		params[i] = p.String()
	}

	out.WriteString("macro")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(m.Body.String())
	out.WriteString("\n}")

	return out.String()
}
