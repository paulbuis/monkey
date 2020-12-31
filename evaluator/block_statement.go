package evaluator

import (
	"monkey/ast"
	"monkey/object"
	"monkey/object/function/environment"
)

// mutually recursive with Eval
func evalBlockStatement(
	block ast.BlockStatement,
	env *environment.Environment,
) object.Object {
	var result object.Object

	count := block.StatementCount()
	for index := 0; index < count; index++ {
		result = Eval(block.Statement(index), env)

		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}
	}

	return result
}
