package hash

import (
	"monkey/object"
	objectError "monkey/object/error"
	objectHash "monkey/object/hash"
	objectNull "monkey/object/null"
	"reflect"
)

func EvalHashIndexExpression(hash, index object.Object) object.Object {
	hashObject, castOk := hash.(*objectHash.Hash)
	if !castOk {
		typeName := reflect.TypeOf(hash).Name()
		return objectError.New2("not a Hash object monkey type:%s golang type: %s", index.Type(), typeName)
	}

	key, ok := index.(object.Hashable)
	if !ok {
		typeName := reflect.TypeOf(hash).Name()
		return objectError.New2("unusable as hash key: %s, golang type: %s", index.Type(), typeName)
	}

	pair, hashLookupSuccess := hashObject.Pairs[key.HashKey()]
	if !hashLookupSuccess {
		return objectNull.NULL()
	}

	return pair.Value
}
