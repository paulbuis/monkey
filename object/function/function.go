package function

import (
	"bytes"
	"monkey/ast"
	"monkey/object"
	"monkey/object/function/environment"
	"strings"
)

type Function struct {
	parameters []ast.Identifier
	body       ast.BlockStatement
	env        *environment.Environment
}

func New(
	parameters []ast.Identifier,
	body ast.BlockStatement,
	env *environment.Environment,
) *Function {
	return &Function{parameters: parameters, body: body, env: env}
}

func (f *Function) Parameters() []ast.Identifier {
	return f.parameters
}

func (f *Function) Body() ast.BlockStatement {
	return f.body
}

func (f *Function) Env() *environment.Environment {
	return f.env
}

func (f *Function) Type() object.ObjectType { return object.FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := make([]string, len(f.parameters))
	for i, p := range f.parameters {
		params[i] = p.String()
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.body.String())
	out.WriteString("\n}")

	return out.String()
}
