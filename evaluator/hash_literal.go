package evaluator

import (
	"monkey/ast"
	"monkey/object"
	objectError "monkey/object/error"
	"monkey/object/function/environment"
	objectHash "monkey/object/hash"
)

func evalHashLiteral(
	node ast.HashLiteral,
	env *environment.Environment,
) object.Object {
	pairs := make(map[object.HashKey]objectHash.Pair)
	count := node.PairCount()
	//values := node.Values()
	// i, keyNode := range node.Keys() {
	for index := 0; index < count; index++ {
		keyNode := node.Keys(index)
		key := Eval(keyNode, env)
		if object.IsError(key) {
			return key
		}

		hashKey, ok := key.(object.Hashable)
		if !ok {
			return objectError.New("unusable as hash key: %s", key.Type())
		}

		valueNode := node.Values(index)
		value := Eval(valueNode, env)
		if object.IsError(value) {
			return value
		}

		hashed := hashKey.HashKey()
		pairs[hashed] = objectHash.Pair{Key: hashKey, Value: value}
	}
	return &objectHash.Hash{Pairs: pairs}
}
