package array

import (
	"bytes"
	"monkey/object"
	objectNull "monkey/object/null"
	"strings"
)

type Array interface {
	Type() object.ObjectType
	Inspect() string
	ElementAt(index int64) object.Object
	Length() int64
}

type Slicer interface {
	Array
	Slice(min, max int64) Array
}

type Pusher interface {
	Array
	Push(newElement object.Object) Array
}

// conforms to interface Array
// conforms to interface Pusher
// conforms to interface Slicer
type Slice struct {
	elements []object.Object
}

func New(elements []object.Object) Array {
	newElements := make([]object.Object, 0, len(elements))
	for _, element := range elements {
		newElements = append(newElements, element)
	}
	return &Slice{elements: elements}
}

func (sl *Slice) Slice(min, max int64) Array {
	return New(sl.elements[min:max])
}

func (sl *Slice) Push(newElement object.Object) Array {
	newElements := make([]object.Object, len(sl.elements), len(sl.elements)+1)
	copy(newElements, sl.elements)
	newElements = append(sl.elements, newElement)
	return &Slice{elements: newElements}
}

func (sl *Slice) ElementAt(index int64) object.Object {
	i := int(index)
	if i < 0 || i >= len(sl.elements) {
		//return objectError.New("index out of bounds index=%d, array length=%d", i, len(ao.elements))
		return objectNull.NULL()
	}
	return sl.elements[index]
}

func (sl *Slice) Length() int64 {
	return int64(len(sl.elements))
}

func (sl *Slice) Type() object.ObjectType {
	return object.ARRAY_OBJ
}

func (sl *Slice) Inspect() string {
	var out bytes.Buffer

	elements := make([]string, 0, len(sl.elements))
	for _, e := range sl.elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
