package boolean

import (
	"fmt"
	"monkey/object"
)

type Boolean struct {
	value bool
}

var (
	TRUE  = &Boolean{value: true}
	FALSE = &Boolean{value: false}
)

func GetBoolean(b bool) *Boolean {
	if b {
		return TRUE
	}
	return FALSE
}

func (b *Boolean) Value() bool {
	return b.value
}

func (b *Boolean) Type() object.ObjectType { return object.BOOLEAN_OBJ }
func (b *Boolean) Inspect() string         { return fmt.Sprintf("%t", b.value) }
func (b *Boolean) HashKey() object.HashKey {
	var value uint64

	if b.value {
		value = 1
	} else {
		value = 0
	}

	return object.HashKey{Type: b.Type(), Value: value}
}
