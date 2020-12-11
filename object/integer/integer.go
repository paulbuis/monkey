package integer

import (
	"fmt"
	"math/big"
	"monkey/object"
)

type Integer struct {
	value *big.Int
}

func New(value int64) *Integer {
	var bigValue big.Int
	newValue := bigValue.SetInt64(value)
	return &Integer{value: newValue}
}

func (i *Integer) Value() (int64, bool) {
	if i.value.IsInt64() {
		return i.value.Int64(), true
	}
	return 0, false
}

func (i *Integer) Type() object.ObjectType { return object.INTEGER_OBJ }
func (i *Integer) Inspect() string         { return fmt.Sprintf("%d", i.value) }
func (i *Integer) HashKey() object.HashKey {
	result := i.value.Uint64()
	return object.HashKey{Type: i.Type(), Value: result} // does not always succeed for large values!!!
}
