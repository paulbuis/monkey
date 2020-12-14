package evaluator

import (
	"monkey/ast"
	"monkey/object"
	objectError "monkey/object/error"
	"monkey/object/function/environment"
	objectHash "monkey/object/hash"
)

func evalHashLiteral(
	node *ast.HashLiteral,
	env *environment.Environment,
) object.Object {
	pairs := make(map[object.HashKey]objectHash.Pair)
	values := node.Values()
	for i, keyNode := range node.Keys() {
		key := Eval(keyNode, env)
		if object.IsError(key) {
			return key
		}

		hashKey, ok := key.(object.Hashable)
		if !ok {
			return objectError.New("unusable as hash key: %s", key.Type())
		}

		valueNode := values[i]
		value := Eval(valueNode, env)
		if object.IsError(value) {
			return value
		}

		hashed := hashKey.HashKey()
		pairs[hashed] = objectHash.Pair{Key: hashKey, Value: value}
	}

	return &objectHash.Hash{Pairs: pairs}
}
