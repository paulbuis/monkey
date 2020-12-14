package evaluator

import (
	evalString "monkey/evaluator/string"
	"monkey/object"
	objectBoolean "monkey/object/boolean"
	objectError "monkey/object/error"
	objectInteger "monkey/object/integer"
)

func applyInfixOperator(
	operator string,
	left, right object.Object,
) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return objectInteger.ApplyIntegerInfixOperator(operator, left, right)

	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalString.ApplyStringInfixOperator(operator, left, right)

	case operator == "==":
		return objectBoolean.GetBoolean(left == right)

	case operator == "!=":
		return objectBoolean.GetBoolean(left != right)

	case left.Type() != right.Type():
		return objectError.New3("type mismatch: %s %s %s",
			left.Type(), operator, right.Type())

	default:
		return objectError.New3("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}
