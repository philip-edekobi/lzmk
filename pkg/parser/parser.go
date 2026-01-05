package parser

import (
	"fmt"

	"github.com/philip-edekobi/lzmk/pkg/lexer"
)

type Parser struct {
	tokens      []*lexer.Token
	pos         int
	MetaHashMap map[string]string
}

func NewParser(tokens []*lexer.Token) *Parser {
	return &Parser{
		tokens:      tokens,
		pos:         0,
		MetaHashMap: make(map[string]string),
	}
}

func (p *Parser) advance() *lexer.Token {
	if p.pos >= len(p.tokens) {
		return &lexer.Token{Kind: lexer.EOF, Value: ""}
	}

	p.pos++
	return p.tokens[p.pos-1]
}

func (p *Parser) peek() *lexer.Token {
	if p.pos >= len(p.tokens) {
		return &lexer.Token{Kind: lexer.EOF, Value: ""}
	}

	return p.tokens[p.pos]
}

func (p *Parser) consumeToken(tokenKind lexer.TokenKind) (*lexer.Token, error) {
	nextToken := p.peek()
	if nextToken.Kind != tokenKind {
		return nil, fmt.Errorf("syntax error on %d:%d\nexpected %v, got %v", nextToken.Line, nextToken.Col, tokenKind, nextToken.Kind)
	}

	return p.advance(), nil
}

func (p *Parser) Parse() (*AST, error) {
	root := newNode(RootNode)
	root.StringValue = "ROOT"
	ast := newAST(root)

	// We should maintain an invariant that the root should have 2 children:
	// first one of Kind TitleNode and the other of Kind BodyNode

	titleNode, err := p.parseTitleNode()
	if err != nil {
		return nil, err
	}

	ast.Root.Children = append(ast.Root.Children, titleNode)

	bodyNode, err := p.parseBody()
	if err != nil {
		return nil, err
	}

	ast.Root.Children = append(ast.Root.Children, bodyNode)

	return ast, nil
}
