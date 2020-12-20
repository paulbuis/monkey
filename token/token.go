package token

import (
	"encoding/json"
)

type TypeName string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT"  // add, foobar, x, y, ...
	INT    = "INT"    // 1343456
	STRING = "STRING" // "foobar"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ    = "=="
	NotEq = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	MACRO    = "MACRO"
)

type Token struct {
	_type   TypeName
	literal string
}

func New(_type TypeName, literal string) Token {
	return Token{_type: _type, literal: literal}
}

func (t Token) Type() TypeName { return t._type }

func (t Token) Literal() string { return t.literal }

var keywords = map[string]TypeName{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"macro":  MACRO,
}

func LookupIdent(ident string) TypeName {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

func (t Token) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["type"] = t._type
	m["literal"] = t.literal
	return json.Marshal(m)
}
