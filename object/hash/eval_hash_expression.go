package hash

import (
	"monkey/object"
	objectError "monkey/object/error"
	objectNull "monkey/object/null"
)

func EvalHashIndexExpression(hash, index object.Object) object.Object {
	hashObject := hash.(*Hash)

	key, ok := index.(object.Hashable)
	if !ok {
		return objectError.New("unusable as hash key: %s", index.Type())
	}

	pair, ok := hashObject.Pairs[key.HashKey()]
	if !ok {
		return objectNull.NULL()
	}

	return pair.Value
}
