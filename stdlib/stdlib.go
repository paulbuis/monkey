package stdlib

// functions in stdlib need injected into the base environment when interpreter starts
// as opposed to builtins whose names are recognized by the parser

import (
	"fmt"
	"monkey/context"
	"monkey/object"
	objectBuiltin "monkey/object/builtin"
	objectNull "monkey/object/null"
)

type StandardLibraryFunction func(ctx *context.Context, args ...object.Object) object.Object

var standardFunctions = map[string]StandardLibraryFunction{
	"puts": puts,
}

func New(ctx *context.Context) map[string]object.Object {
	library := make(map[string]object.Object)
	for name, fun := range standardFunctions {
		library[name] = ConvertToBuiltin(fun, ctx)
	}
	return library
}

func puts(ctx *context.Context, args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Fprintln(ctx.Out(), arg.Inspect())
	}

	return objectNull.NULL()
}

func contextWrapper(ctx *context.Context, f func(ctx *context.Context, args ...object.Object) object.Object) func(args ...object.Object) object.Object {
	return func(args ...object.Object) object.Object {
		return f(ctx, args...)
	}
}

func ConvertToBuiltin(f StandardLibraryFunction, ctx *context.Context) *objectBuiltin.Builtin {
	return objectBuiltin.New(contextWrapper(ctx, f))
}
