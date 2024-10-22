package parser

import (
	"Alastair7/monke/ast"
	"Alastair7/monke/lexer"
	"Alastair7/monke/token"
)

type Parser struct {
	l *lexer.Lexer

	currentToken token.Token
	peekToken token.Token
}

func New(l* lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Set currentToken and peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken;
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}