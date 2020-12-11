package return_value

import (
	"monkey/object"
)

type ReturnValue struct {
	value object.Object
}

func New(value object.Object) *ReturnValue {
	return &ReturnValue{value: value}
}

func (rv *ReturnValue) Value() object.Object {
	return rv.value
}

func (rv *ReturnValue) Type() object.ObjectType {
	return object.RETURN_VALUE_OBJ
}

func (rv *ReturnValue) Inspect() string {
	return rv.value.Inspect()
}
