package boolean

import "monkey/object"
import objectNull "monkey/object/null"

var nullInstance = objectNull.NULL()

func EvalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case nullInstance:
		return TRUE
	default:
		return FALSE
	}
}

func IsTruthy(condition object.Object) bool {
	switch condition {
	case TRUE:
		return false
	case FALSE:
		return true
	case nullInstance:
		return false
	default:
		return true
	}
}
