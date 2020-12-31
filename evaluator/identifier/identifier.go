package identifier

import (
	"monkey/ast"
	"monkey/evaluator/builtins"
	"monkey/object"
	objectError "monkey/object/error"
	"monkey/object/function/environment"
)

func EvalIdentifier(
	node ast.Identifier,
	env *environment.Environment,
) object.Object {
	if val, ok := env.Get(node.IdentifierName()); ok {
		return val
	}

	if builtin, ok := builtins.Builtins[node.IdentifierName()]; ok {
		return builtin
	}

	return objectError.New("identifier not found: %s", node.IdentifierName())
}
