package builtins

import (
	"monkey/object"
	objectArray "monkey/object/array"
	"monkey/object/boolean"
	objectBuiltin "monkey/object/builtin"
	objectError "monkey/object/error"
	objectInteger "monkey/object/integer"
	objectNull "monkey/object/null"
	objectString "monkey/object/string"
	"reflect"
)

func builtinLen(args ...object.Object) object.Object {
	if len(args) != 1 {
		return objectError.New("wrong number of arguments. got=%d, want=1",
			len(args))
	}

	switch arg := args[0].(type) {
	case objectArray.Array:
		n := arg.Length()
		return objectInteger.New(n)
	case *objectString.String:
		n := len(arg.Value())
		return objectInteger.New(int64(n))
	default:
		typeName := reflect.TypeOf(args[0]).Name()
		return objectError.New2("argument to `len` not supported, got %s, golang type: %v",
			args[0].Type(), typeName)
	}
}

func isArray(args ...object.Object) object.Object {
	if len(args) != 1 {
		return objectError.New("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	return boolean.GetBoolean(object.IsArray(args[0]))
}

func isBoolean(args ...object.Object) object.Object {
	if len(args) != 1 {
		return objectError.New("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	return boolean.GetBoolean(object.IsBoolean(args[0]))
}

func isError(args ...object.Object) object.Object {
	if len(args) != 1 {
		return objectError.New("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	return boolean.GetBoolean(object.IsError(args[0]))
}

func isInteger(args ...object.Object) object.Object {
	if len(args) != 1 {
		return objectError.New("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	return boolean.GetBoolean(object.IsInteger(args[0]))
}

func isFunction(args ...object.Object) object.Object {
	if len(args) != 1 {
		return objectError.New("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	return boolean.GetBoolean(object.IsFunction(args[0]))
}

func isHash(args ...object.Object) object.Object {
	if len(args) != 1 {
		return objectError.New("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	return boolean.GetBoolean(object.IsHash(args[0]))
}

func isString(args ...object.Object) object.Object {
	if len(args) != 1 {
		return objectError.New("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	return boolean.GetBoolean(object.IsString(args[0]))
}

func typeof(args ...object.Object) object.Object {
	if len(args) != 1 {
		return objectError.New("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	return objectString.New(string(args[0].Type()))
}

var Builtins = map[string]*objectBuiltin.Builtin{
	"len":        objectBuiltin.New(builtinLen),
	"isArray":    objectBuiltin.New(isArray),
	"isBoolean":  objectBuiltin.New(isBoolean),
	"isError":    objectBuiltin.New(isError),
	"isInteger":  objectBuiltin.New(isInteger),
	"isHash":     objectBuiltin.New(isHash),
	"isString":   objectBuiltin.New(isString),
	"isFunction": objectBuiltin.New(isFunction),
	"typeof":     objectBuiltin.New(typeof),

	"first": objectBuiltin.New(
		func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return objectError.New("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return objectError.New("argument to `first` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(objectArray.Array)
			if arr.Length() > 0 {
				return arr.ElementAt(0)
			}

			return objectNull.NULL() // or should this be an error???
		},
	),
	"last": objectBuiltin.New(
		func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return objectError.New("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return objectError.New("argument to `last` must be ARRAY, got %s",
					args[0].Type())
			}

			arr, ok := args[0].(objectArray.Array)
			if !ok {
				typeName := reflect.TypeOf(args[0]).Name()
				return objectError.New("argument to slice does not conform to Slicer`, got go type %s", typeName)
			}
			length := arr.Length()
			return arr.ElementAt(length - 1)
		},
	),
	"rest": objectBuiltin.New(
		func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return objectError.New("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return objectError.New("argument to `slice` must be ARRAY, got %s",
					args[0].Type())
			}

			arr, sliceable := args[0].(objectArray.Slicer)
			if !sliceable {
				typeName := reflect.TypeOf(args[0]).Name()
				return objectError.New("argument to slice does not conform to Slicer`, got go type %s", typeName)
			}
			length := arr.Length()
			if length > 0 {
				return arr.Slice(1, length)
			}

			return objectNull.NULL() // or Error???
		},
	),
	"push": objectBuiltin.New(
		func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return objectError.New("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return objectError.New("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}

			pusher, pusherCheck := args[0].(objectArray.Pusher)
			if !pusherCheck {
				typeName := reflect.TypeOf(args[0]).Name()
				return objectError.New("argument to slice does not conform to Slicer`, got go type %s", typeName)
			}
			return pusher.Push(args[1])
		},
	),
}
