package hash

import (
	"bytes"
	"fmt"
	"monkey/object"
	objectError "monkey/object/error"
	"strings"
)

type Pair struct {
	Key   object.Object
	Value object.Object
}

type Hash struct {
	Pairs map[object.HashKey]Pair
}

func NewHash(keys []object.Object, values []object.Object) object.Object {
	pairMap := make(map[object.HashKey]Pair)
	for i, key := range keys {
		hashableKey, ok := key.(object.Hashable)
		if !ok {
			return objectError.New("unhashable key:", key.Inspect())
		}
		hashKey := hashableKey.HashKey()
		value := values[i]
		pair := Pair{Key: key, Value: value}
		pairMap[hashKey] = pair
	}
	return &Hash{Pairs: pairMap}
}

func (h *Hash) Type() object.ObjectType { return object.HASH_OBJ }
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := make([]string, 0, len(h.Pairs))
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
