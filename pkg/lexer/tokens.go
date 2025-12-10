package lexer

import "fmt"

type TokenKind int8

const (
	_ TokenKind = iota
	TitleHash
	HeaderHash
	MetaHash
	HashBang
	String
	LeftParen
	RightParen
	LeftBrace
	RightBrace
	NewLine
	EOF
)

type Token struct {
	kind  TokenKind
	value string
	line  int
	col   int
}

func (t *Token) String() string {
	return fmt.Sprintf("TokenKind %v: %s at position (%d, %d)", t.kind, t.value, t.line, t.col)
}
