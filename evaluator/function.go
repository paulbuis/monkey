package evaluator

import (
	"monkey/object"
	objectBuiltin "monkey/object/builtin"
	objectError "monkey/object/error"
	objectFunction "monkey/object/function"
	"monkey/object/function/environment"
	objectReturnValue "monkey/object/return_value"
)

// mutually recursive with Eval
func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {

	case *objectFunction.Function:
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body(), extendedEnv)
		return unwrapReturnValue(evaluated)

	case *objectBuiltin.Builtin:
		return fn.Fn()(args...)

	default:
		return objectError.New("not a function: %s", fn.Type())
	}
}

// only called by applyFunction
func extendFunctionEnv(
	fn *objectFunction.Function,
	args []object.Object,
) *environment.Environment {
	env := environment.NewEnclosedEnvironment(fn.Env())

	for paramIdx, param := range fn.Parameters() {
		env.Set(param.Value(), args[paramIdx])
	}

	return env
}

// only called by applyFunction
func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*objectReturnValue.ReturnValue); ok {
		return returnValue.Value()
	}

	return obj
}
