package string

import (
	"monkey/object"
	object_error "monkey/object/error"
)

func EvalStringInfixExpression(
	operator string,
	left, right object.Object) object.Object {

	if operator != "+" {
		return object_error.New("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}

	leftVal := left.(*String).value
	rightVal := right.(*String).value
	return New(leftVal + rightVal)
}
