package lexer

import (
	"unicode"
)

type Lexer struct {
	source  []rune
	start   int
	pos     int
	col     int
	line    int
	inParen bool
	inBrace bool
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		source:  []rune(input),
		start:   0,
		pos:     0,
		line:    1,
		col:     1,
		inParen: false,
		inBrace: false,
	}
}

func (l *Lexer) advance() rune {
	if l.pos >= len(l.source) {
		return 0
	}

	l.pos++
	l.col++
	return l.source[l.pos-1]
}

func (l *Lexer) peek() rune {
	if l.pos >= len(l.source) {
		return 0
	}

	return l.source[l.pos]
}

func (l *Lexer) Lex() ([]*Token, error) {
	tokens := []*Token{}

	for l.pos < len(l.source) {
		char := l.advance()
		for unicode.IsSpace(char) {
			if char == '\n' {
				l.line++
				col := l.col
				l.col = 1

				tokens = append(
					tokens,
					&Token{kind: NewLine, value: "\n", line: l.line - 1, col: col - 1},
				)
			}
			char = l.advance()
		}

		switch char {
		case '#':
			if l.peek() == '#' {
				l.advance()

				if l.peek() == '#' {
					l.advance()

					tokens = append(
						tokens,
						&Token{kind: MetaHash, value: "###", line: l.line, col: l.col - 3},
					)
				} else {
					l.pos++
					tokens = append(tokens, &Token{kind: HeaderHash, value: "##", line: l.line, col: l.col - 2})
				}
			} else if l.peek() == '!' {
				l.advance()

				tokens = append(
					tokens,
					&Token{kind: HashBang, value: "#!", line: l.line, col: l.col - 2},
				)
			} else {
				tokens = append(tokens, &Token{kind: TitleHash, value: "#", line: l.line, col: l.col - 1})
			}

		case '(':
			l.inParen = true

			tokens = append(
				tokens,
				&Token{kind: LeftParen, value: "(", line: l.line, col: l.col - 1},
			)

		case ')':
			l.inParen = false

			tokens = append(
				tokens,
				&Token{kind: RightParen, value: ")", line: l.line, col: l.col - 1},
			)

		case '[':
			l.inBrace = true

			tokens = append(
				tokens,
				&Token{kind: LeftBrace, value: "[", line: l.line, col: l.col - 1},
			)

		case ']':
			l.inBrace = false

			tokens = append(
				tokens,
				&Token{kind: RightBrace, value: "]", line: l.line, col: l.col - 1},
			)

		case 0:
			tokens = append(tokens, &Token{kind: EOF, value: "", line: l.line, col: l.col - 1})

		default:
			start := l.col
			s := string(char)

			for {
				if l.peek() == 0 || l.peek() == '\n' || (l.peek() == ')' && l.inParen) ||
					(l.peek() == ']' && l.inBrace) {
					break
				}

				char = l.advance()
				s += string(char)
			}

			tokens = append(tokens, &Token{kind: String, value: s, line: l.line, col: start - 1})
		}
	}

	return tokens, nil
}
