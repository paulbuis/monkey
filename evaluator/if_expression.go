package evaluator

import (
	"monkey/ast"
	"monkey/evaluator/boolean"
	"monkey/object"
	"monkey/object/function/environment"
	objectNull "monkey/object/null"
)

func evalIfExpression(
	ie ast.IfExpression,
	env *environment.Environment,
) object.Object {
	condition := Eval(ie.Condition(), env)
	if object.IsError(condition) {
		return condition
	}

	if boolean.IsTruthy(condition) {
		return Eval(ie.Consequence(), env)
	} else if ie.Alternative() != nil {
		return Eval(ie.Alternative(), env)
	} else {
		return objectNull.NULL()
	}
}
