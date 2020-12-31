package ast

import (
	"bytes"
	"monkey/ast"
	"monkey/token"
)

//  *BlockStatement conforms to interface ast.BlockStatement
type BlockStatement struct {
	token      token.Token // the { token
	statements []ast.Statement
}

// diagnostic check to verify *BlockStatement struct
// in this package conforms to ast.BlockStatement interface
var _ ast.BlockStatement = &BlockStatement{}

func NewBlockStatement(token token.Token, statements []ast.Statement) *BlockStatement {
	return &BlockStatement{token: token, statements: statements}
}

func (bs *BlockStatement) Token() token.Token {
	return bs.token
}

func (bs *BlockStatement) StatementNode() {}

func (bs *BlockStatement) Statement(index int) ast.Statement {
	return bs.statements[index]
}

func (bs *BlockStatement) StatementCount() int {
	return len(bs.statements)
}

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
