package boolean

import (
	"monkey/object"
	"monkey/object/boolean"
	objectNull "monkey/object/null"
)

var nullInstance = objectNull.NULL()

func ApplyBangOperator(right object.Object) object.Object {
	switch right {
	case boolean.TRUE:
		return boolean.FALSE
	case boolean.FALSE:
		return boolean.TRUE
	case nullInstance:
		return boolean.TRUE
	default:
		return boolean.FALSE
	}
}

func IsTruthy(condition object.Object) bool {
	switch condition {
	case boolean.TRUE:
		return true
	case boolean.FALSE:
		return false
	case nullInstance:
		return false
	default:
		return true
	}
}
