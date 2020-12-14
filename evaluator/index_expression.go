package evaluator

import (
	"monkey/evaluator/array"
	evalHash "monkey/evaluator/hash"
	"monkey/object"
	objectError "monkey/object/error"
)

func applyIndexOperator(left, index object.Object) object.Object {
	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return array.EvalArrayIndexExpression(left, index)
	case left.Type() == object.HASH_OBJ:
		return evalHash.EvalHashIndexExpression(left, index)
	default:
		return objectError.New("index operator not supported: %s", left.Type())
	}
}
