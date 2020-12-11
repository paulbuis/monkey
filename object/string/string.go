package string

import (
	"hash/fnv"
	"monkey/object"
)

type String struct {
	value string
}

func New(value string) *String {
	return &String{value: value}
}

func (s *String) Value() string {
	return s.value
}

func (s *String) Type() object.ObjectType {
	return object.STRING_OBJ
}

func (s *String) Inspect() string {
	return s.value
}

func (s *String) HashKey() object.HashKey {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s.value)) // not handling errors

	return object.HashKey{Type: s.Type(), Value: h.Sum64()}
}
