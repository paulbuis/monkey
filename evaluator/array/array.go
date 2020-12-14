package array

import (
	"monkey/object"
	objectArray "monkey/object/array"
	objectError "monkey/object/error"
	objectInteger "monkey/object/integer"
)

// returns an error if array does not conform to Array interface
// behaves as if index was -1 if index does not conform to Integer interface
func EvalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObject, okCast := array.(objectArray.Array)
	if !okCast {
		return objectError.New("index operator not supported: %s", array.Type())
	}
	idx, ok := index.(*objectInteger.Integer).Value()
	if !ok {
		idx = -1
	}

	return arrayObject.ElementAt(idx)
}
