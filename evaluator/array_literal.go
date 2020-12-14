package evaluator

import (
	"monkey/ast"
	"monkey/object"
	objectArray "monkey/object/array"
	"monkey/object/function/environment"
)

func evalArrayLiteral(arrayLiteral *ast.ArrayLiteral, env *environment.Environment) object.Object {
	elements := evalExpressions(arrayLiteral.Elements(), env)
	if len(elements) == 1 && object.IsError(elements[0]) {
		return elements[0]
	}
	return objectArray.New(elements)
}
