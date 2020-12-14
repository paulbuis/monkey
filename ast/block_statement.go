package ast

import (
	"bytes"
	"fmt"
	"monkey/token"
)

// represents a sequence of statements
//   conforms to interface Node
//   conforms to interface Statement
//   conforms to interface fmt.Stringer
type BlockStatement struct {
	token      token.Token // the { token
	statements []Statement
}

func NewBlockStatement(token token.Token, statements []Statement) *BlockStatement {
	return &BlockStatement{token: token, statements: statements}
}

func (bs *BlockStatement) Token() token.Token {
	return bs.token
}

func (bs *BlockStatement) Statements() []Statement {
	return bs.statements
}

func (bs *BlockStatement) MarshalJSON() ([]byte, error) {
	var out bytes.Buffer
	var err error
	out.WriteString(`{"NodeType":"BlockStatement"`)
	out.WriteString(`," token": `)
	buf, _ := bs.token.MarshalJSON()
	out.Write(buf)
	out.WriteString(`, "statements": [`)

	fmt.Println("Marshalling a BlockStatement")
	for i, s := range bs.statements {
		fmt.Println("s =", s.String())
		buf, err = s.MarshalJSON()
		if err != nil {
			return []byte{}, err
		}
		out.Write(buf)
		if i != len(bs.statements)-1 {
			out.WriteRune(',')
		}
	}

	out.WriteString("]}")
	return out.Bytes(), nil
}
func (bs *BlockStatement) statementNode() {}

func (bs *BlockStatement) TokenLiteral() string {
	return bs.token.Literal()
}

func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	out.WriteRune('{')

	for _, s := range bs.statements {
		out.WriteString(s.String())
	}

	out.WriteRune('}')

	return out.String()
}
