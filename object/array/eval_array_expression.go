package array

import (
	"monkey/object"
	objectInteger "monkey/object/integer"
)

func EvalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObject, _ := array.(Array)
	idx, ok := index.(*objectInteger.Integer).Value()
	if !ok {
		idx = -1
	}

	return arrayObject.ElementAt(idx)
}
