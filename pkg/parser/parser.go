package parser

import "github.com/philip-edekobi/lzmk/pkg/lexer"

type Parser struct {
	tokens []*lexer.Token
	pos    int
}

func NewParser(tokens []*lexer.Token) *Parser {
	return &Parser{
		tokens: tokens,
		pos:    0,
	}
}

func (p *Parser) advance() *lexer.Token {
	if p.pos >= len(p.tokens) {
		return &lexer.Token{Kind: lexer.EOF, Value: ""}
	}

	p.pos++
	return p.tokens[p.pos-1]
}

func (p *Parser) consumeToken(tokenKind lexer.TokenKind) (*lexer.Token, error) {
	nextToken := p.peek()
}

func (p *Parser) peek() *lexer.Token {
	if p.pos >= len(p.tokens) {
		return &lexer.Token{Kind: lexer.EOF, Value: ""}
	}

	return p.tokens[p.pos]
}

func (p *Parser) Parse(input []*lexer.Token) (*AST, error) {
	ast := &AST{}

	return ast, nil
}
