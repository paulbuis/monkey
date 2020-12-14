package ast

import (
	"bytes"
	"monkey/token"
)

// conforms to interface Node
// conforms to interface Expression
// conforms to interface fmt.Stringer
type IfExpression struct {
	token       token.Token // The 'if' token
	condition   Expression
	consequence *BlockStatement
	alternative *BlockStatement
}

func NewIfExpression(token token.Token,
	condition Expression,
	consequence *BlockStatement,
	alternative *BlockStatement) *IfExpression {
	return &IfExpression{token: token, condition: condition, consequence: consequence, alternative: alternative}
}

func (ie *IfExpression) Token() token.Token {
	return ie.token
}

func (ie *IfExpression) Condition() Expression {
	return ie.condition
}

func (ie *IfExpression) Consequence() *BlockStatement {
	return ie.consequence
}

func (ie *IfExpression) Alternative() *BlockStatement {
	return ie.alternative
}

func (ie *IfExpression) MarshalJSON() ([]byte, error) {
	var out bytes.Buffer
	var err error
	out.WriteString(`{"NodeType":"IfExpression",`)

	out.WriteString(`" token":`)
	buf, _ := ie.token.MarshalJSON()
	out.Write(buf)

	out.WriteString(`,"condition":`)
	buf, err = ie.condition.MarshalJSON()
	if err != nil {
		return []byte{}, err
	}
	out.Write(buf)

	out.WriteString(`,"consequence":`)
	buf, err = ie.consequence.MarshalJSON()
	if err != nil {
		return []byte{}, err
	}
	out.Write(buf)

	if ie.alternative != nil {
		out.WriteString(`,"else_part":`)
		buf, err = ie.alternative.MarshalJSON()
		if err != nil {
			return []byte{}, err
		}
		out.Write(buf)
	}
	out.WriteString("}")
	return out.Bytes(), nil
}

func (ie *IfExpression) expressionNode() {}

func (ie *IfExpression) TokenLiteral() string {
	return ie.token.Literal()
}

func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if (")
	out.WriteString(ie.condition.String())
	out.WriteString(") ")
	out.WriteString(ie.consequence.String())

	if ie.alternative != nil {
		out.WriteString(" else ")
		out.WriteString(ie.alternative.String())
	}

	return out.String()
}
