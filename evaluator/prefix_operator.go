package evaluator

import (
	"monkey/evaluator/boolean"
	"monkey/object"
	objectError "monkey/object/error"
	objectInteger "monkey/object/integer"
)

func applyPrefixOperator(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return boolean.ApplyBangOperator(right)
	case "-":
		return objectInteger.ApplyMinusPrefixOperator(right)
	default:
		return objectError.New2("unknown operator: %s%s", operator, right.Type())
	}
}
