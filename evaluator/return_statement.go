package evaluator

import (
	"monkey/ast"
	"monkey/object"
	"monkey/object/function/environment"
	objectReturnValue "monkey/object/return_value"
)

// mutually recursive with Eval
func evalReturnStatement(
	node ast.ReturnStatement,
	env *environment.Environment,
) object.Object {
	r := node.ReturnValue()
	val := Eval(r, env)

	if object.IsError(val) {
		return val
	}

	return objectReturnValue.New(val)
}
