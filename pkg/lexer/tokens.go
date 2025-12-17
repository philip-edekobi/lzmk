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
	Kind  TokenKind
	Value string
	Line  int
	Col   int
}

func (t *Token) String() string {
	return fmt.Sprintf("TokenKind %v: %s at position (%d, %d)", t.Kind, t.Value, t.Line, t.Col)
}
