package ast

import (
	"bytes"
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

func (bs *BlockStatement) statementNode() {}

func (bs *BlockStatement) TokenLiteral() string {
	return bs.token.Literal()
}

func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.statements {
		out.WriteString(s.String())
	}

	return out.String()
}
