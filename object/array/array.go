package array

import (
	"bytes"
	"monkey/object"
	objectError "monkey/object/error"
	"strings"
)

type Array interface {
	Type() object.ObjectType
	Inspect() string
	ElementAt(index int64) object.Object
	Length() int64
	Slice(min, max int64) Array
	Push(newElement object.Object) Array
}

type slice struct {
	elements []object.Object
}

func New(elements []object.Object) Array {
	newElements := make([]object.Object, 0, len(elements))
	for _, element := range elements {
		newElements = append(newElements, element)
	}
	return &slice{elements: elements}
}

func (ao *slice) Slice(min, max int64) Array {
	return New(ao.elements[min:max])
}

func (ao *slice) Push(newElement object.Object) Array {
	newElements := make([]object.Object, len(ao.elements), len(ao.elements)+1)
	copy(newElements, ao.elements)
	newElements = append(ao.elements, newElement)
	return &slice{elements: newElements}

}

func (ao *slice) ElementAt(index int64) object.Object {
	i := int(index)
	if i < 0 || i >= len(ao.elements) {
		return objectError.New("index out of bounds index=%d, array length=%d", i, len(ao.elements))
	}
	return ao.elements[index]
}

func (ao *slice) Length() int64 {
	return int64(len(ao.elements))
}

func (ao *slice) Type() object.ObjectType {
	return object.ARRAY_OBJ
}

func (ao *slice) Inspect() string {
	var out bytes.Buffer

	elements := make([]string, 0, len(ao.elements))
	for _, e := range ao.elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
