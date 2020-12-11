package quote

import (
	"monkey/ast"
	"monkey/object"
)

type Quote struct {
	Node ast.Node
}

func (q *Quote) Type() object.ObjectType {
	return object.QUOTE_OBJ
}

func (q *Quote) Inspect() string {
	return "QUOTE(" + q.Node.String() + ")"
}
