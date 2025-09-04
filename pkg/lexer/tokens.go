package lexer

import "fmt"

type TokenKind int8

const (
	_ TokenKind = iota
	TitleHash
	HeaderHash
	MetaHash
	HashBang
	NarrowString
	MultiString
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
	row   int
	col   int
}

func (t *Token) String() string {
	return fmt.Sprintf("TokenKind %v: %s at position (%d, %d)", t.kind, t.value, t.row, t.col)
}
