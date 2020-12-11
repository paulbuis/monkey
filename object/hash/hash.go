package hash

import (
	"bytes"
	"fmt"
	"monkey/object"
	"strings"
)

type HashPair struct {
	Key   object.Object
	Value object.Object
}

type Hash struct {
	Pairs map[object.HashKey]HashPair
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
