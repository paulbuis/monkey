package string

import (
	"monkey/object"
	objectError "monkey/object/error"
	objectString "monkey/object/string"
)

func ApplyStringInfixOperator(
	operator string,
	left, right object.Object) object.Object {

	if operator != "+" {
		return objectError.New3("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}

	leftVal := left.(*objectString.String).Value()
	rightVal := right.(*objectString.String).Value()
	return objectString.New(leftVal + rightVal)
}
