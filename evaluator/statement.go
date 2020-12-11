package evaluator

import (
	"monkey/ast"
	"monkey/object"
	objectReturnValue "monkey/object/return_value"
)

func evalBlockStatement(
	block *ast.BlockStatement,
	env *object.Environment,
) object.Object {
	var result object.Object

	for _, statement := range block.Statements() {
		result = Eval(statement, env)

		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}
	}

	return result
}

func evalReturnStatement(
	node ast.Node,
	env *object.Environment,
) object.Object {
	val := Eval(node, env)
	if isError(val) {
		return val
	}
	return objectReturnValue.New(val)
}
