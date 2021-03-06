package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	parserAst "monkey/parser/ast"
	"monkey/token"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
	INDEX       // array[index]
)

var precedences = map[token.TypeName]int{
	token.EQ:       EQUALS,
	token.NotEq:    EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
	token.LBRACKET: INDEX,
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TypeName]prefixParseFn
	infixParseFns  map[token.TypeName]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[token.TypeName]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.STRING, p.parseStringLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.IF, p.parseIfExpression)
	p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)
	p.registerPrefix(token.LBRACKET, p.parseArrayLiteral)
	p.registerPrefix(token.LBRACE, p.parseHashLiteral)

	p.infixParseFns = make(map[token.TypeName]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NotEq, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)

	p.registerInfix(token.LPAREN, p.parseCallExpression)
	p.registerInfix(token.LBRACKET, p.parseIndexExpression)

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.TypeName) bool {
	return p.curToken.Type() == t
}

func (p *Parser) peekTokenIs(t token.TypeName) bool {
	return p.peekToken.Type() == t
}

func (p *Parser) expectPeek(t token.TypeName) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TypeName) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type())
	p.errors = append(p.errors, msg)
}

func (p *Parser) noPrefixParseFnError(t token.TypeName) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) ParseProgram() ast.Program {
	statements := make([]ast.Statement, 0)

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			statements = append(statements, stmt)
		}
		p.nextToken()
	}

	return parserAst.NewProgram(statements)
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type() {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() ast.LetStatement {
	startToken := p.curToken

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	name := parserAst.NewIdentifier(p.curToken, p.curToken.Literal())

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	value := p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return parserAst.NewLetStatement(startToken, name, value)
}

func (p *Parser) parseReturnStatement() ast.ReturnStatement {
	startToken := p.curToken

	p.nextToken()

	returnValue := p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return parserAst.NewReturnStatement(startToken, returnValue)
}

func (p *Parser) parseExpressionStatement() ast.ExpressionStatement {
	startToken := p.curToken

	expression := p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return parserAst.NewExpressionStatement(startToken, expression)
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type()]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type())
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type()]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type()]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type()]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) parseIdentifier() ast.Expression {
	return parserAst.NewIdentifier(p.curToken, p.curToken.Literal())
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	startToken := p.curToken

	value, err := strconv.ParseInt(p.curToken.Literal(), 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal())
		p.errors = append(p.errors, msg)
		return nil
	}

	return parserAst.NewIntegerLiteral(startToken, value)
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return parserAst.NewStringLiteral(p.curToken, p.curToken.Literal())
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	startToken := p.curToken
	p.nextToken()
	right := p.parseExpression(PREFIX)

	return parserAst.NewPrefixExpression(startToken, startToken.Literal(), right)
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	startToken := p.curToken
	precedence := p.curPrecedence()
	p.nextToken()
	right := p.parseExpression(precedence)

	return parserAst.NewInfixExpression(startToken, left, startToken.Literal(), right)
}

func (p *Parser) parseBoolean() ast.Expression {
	return parserAst.NewBoolean(p.curToken, p.curTokenIs(token.TRUE))
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

func (p *Parser) parseIfExpression() ast.Expression {
	startToken := p.curToken
	//expression := &ast.IfExpression{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	p.nextToken()
	condition := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	consequence := p.parseBlockStatement()
	var alternative ast.BlockStatement
	if p.peekTokenIs(token.ELSE) {
		p.nextToken()

		if !p.expectPeek(token.LBRACE) {
			return nil
		}

		alternative = p.parseBlockStatement()
	}

	return parserAst.NewIfExpression(startToken, condition, consequence, alternative)
}

func (p *Parser) parseBlockStatement() ast.BlockStatement {
	startToken := p.curToken
	statements := make([]ast.Statement, 0)

	p.nextToken()

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			statements = append(statements, stmt)
		}
		p.nextToken()
	}

	return parserAst.NewBlockStatement(startToken, statements)
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	startToken := p.curToken

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	parameters := p.parseFunctionParameters()

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	body := p.parseBlockStatement()

	return parserAst.NewFunctionLiteral(startToken, parameters, body)
}

func (p *Parser) parseFunctionParameters() []ast.Identifier {
	identifiers := make([]ast.Identifier, 0)

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	ident := parserAst.NewIdentifier(p.curToken, p.curToken.Literal())
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := parserAst.NewIdentifier(p.curToken, p.curToken.Literal())
		identifiers = append(identifiers, ident)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return identifiers
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	startToken := p.curToken
	arguments := p.parseExpressionList(token.RPAREN)
	return parserAst.NewCallExpression(startToken, function, arguments)
}

func (p *Parser) parseExpressionList(end token.TypeName) []ast.Expression {
	list := make([]ast.Expression, 0)

	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()
	list = append(list, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(end) {
		return nil
	}

	return list
}

func (p *Parser) parseArrayLiteral() ast.Expression {
	startToken := p.curToken
	elements := p.parseExpressionList(token.RBRACKET)

	return parserAst.NewArrayLiteral(startToken, elements)
}

func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	startToken := p.curToken
	//exp := &ast.IndexExpression{Token: p.curToken, Left: left}

	p.nextToken()
	index := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RBRACKET) {
		return nil
	}

	return parserAst.NewIndexExpression(startToken, left, index)
}

func (p *Parser) parseHashLiteral() ast.Expression {
	startToken := p.curToken
	keys := make([]ast.Expression, 0)
	values := make([]ast.Expression, 0)

	for !p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		keys = append(keys, p.parseExpression(LOWEST))

		if !p.expectPeek(token.COLON) {
			return nil
		}

		p.nextToken()
		value := p.parseExpression(LOWEST)
		values = append(values, value)

		if !p.peekTokenIs(token.RBRACE) && !p.expectPeek(token.COMMA) {
			return nil
		}
	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return parserAst.NewHashLiteral(startToken, keys, values)
}

func (p *Parser) registerPrefix(tokenType token.TypeName, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TypeName, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
