package null

import (
	"monkey/object"
)

type Null struct{}

var null = &Null{}

func NULL() object.Object {
	return null
}

func (n *Null) Type() object.ObjectType {
	return object.NULL_OBJ
}

func (n *Null) Inspect() string {
	return "null"
}
