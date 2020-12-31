package evaluator

import (
	"monkey/object"
	objectBuiltin "monkey/object/builtin"
	objectError "monkey/object/error"
	objectFunction "monkey/object/function"
	"monkey/object/function/environment"
	objectReturnValue "monkey/object/return_value"
)

func wrapFunctionAsThunk(fn *objectFunction.Function,
	args []object.Object,
) func() object.Object {
	return func() object.Object {
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body(), extendedEnv)
		return unwrapReturnValue(evaluated)
	}
}

func wrapBuiltinAsThunk(
	builtin *objectBuiltin.Builtin,
	args []object.Object,
) func() object.Object {
	return func() object.Object {
		return builtin.Fn()(args...)
	}
}

func wrapErrorAsThunk(obj object.Object) func() object.Object {
	return func() object.Object {
		return objectError.New("not a function: %s", obj.Type())
	}
}

// mutually indirectly recursive with Eval
func convertToThunk(fn object.Object,
	args []object.Object,
) func() object.Object {
	switch fn := fn.(type) {

	case *objectFunction.Function:
		return wrapFunctionAsThunk(fn, args)

	case *objectBuiltin.Builtin:
		return wrapBuiltinAsThunk(fn, args)

	default:
		return wrapErrorAsThunk(fn)
	}
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	thunk := convertToThunk(fn, args)
	return thunk()
}

// only called by wrapFunctionAsThunk
func extendFunctionEnv(
	fn *objectFunction.Function,
	args []object.Object,
) *environment.Environment {
	env := environment.NewEnclosedEnvironment(fn.Env())

	for paramIdx, param := range fn.Parameters() {
		env.Set(param.IdentifierName(), args[paramIdx])
	}

	return env
}

// only called by wrapFunctionAsThunk
func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*objectReturnValue.ReturnValue); ok {
		return returnValue.Value()
	}

	return obj
}
