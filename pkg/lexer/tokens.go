package lexer

type TokenKind int8

const (
	_                       = iota
	LiteralString TokenKind = 1 << (2 * iota)
)

type Token struct {
	kind  TokenKind
	value any
	row   int
	col   int
}
