package builtin

import (
	"monkey/object"
)

type Builtin struct {
	fn object.BuiltinFunction
}

func New(fn object.BuiltinFunction) *Builtin {
	return &Builtin{fn: fn}
}

func (b *Builtin) Fn() object.BuiltinFunction {
	return b.fn
}

func (b *Builtin) Type() object.ObjectType {
	return object.BUILTIN_OBJ
}

func (b *Builtin) Inspect() string {
	return "builtin function"
}
