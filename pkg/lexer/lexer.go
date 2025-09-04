package lexer

import (
	"unicode"
)

type Mode int8

const (
	MultiMode Mode = iota
	NarrowMode
	UrlMode
	ImageMultiMode
)

type Lexer struct {
	input     []rune
	lineCount int
	pos       int
	col       int
	prevCol   int
	mode      Mode
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:     []rune(input),
		lineCount: 1,
		col:       1,
		prevCol:   1,
		pos:       0,
	}
}

func (l *Lexer) advance() rune {
	if l.pos >= len(l.input) {
		return 0
	}

	l.prevCol = l.col

	char := l.input[l.pos]
	if char == '\n' {
		l.lineCount++
		l.col = 1
	} else {
		l.col++
	}

	l.pos++
	return char
}

func (l *Lexer) peek() rune {
	if l.pos > len(l.input) {
		return 0
	}

	if l.pos == 0 {
		return l.input[0]
	}

	return l.input[l.pos-1]
}

func (l *Lexer) peekNext() rune {
	if l.pos+1 >= len(l.input) {
		return 0
	}

	return l.input[l.pos]
}

func (l *Lexer) eatWhiteSpace() {
	for l.pos < len(l.input) {
		currentChar := l.peek()

		if currentChar == '\n' || !unicode.IsSpace(currentChar) {
			break
		}
		l.advance()
	}
}

func (l *Lexer) lexNarrowString() *Token {
	val := ""
	row := l.lineCount
	char := l.peek()
	col := l.prevCol

	for char != 0 && !unicode.IsSpace(char) && char != '\n' {
		val += string(char)

		char = l.advance()
	}

	if val == "" {
		return nil
	}

	return &Token{kind: NarrowString, value: val, row: row, col: col}
}

func (l *Lexer) lexMultiString() *Token {
	val := ""
	row := l.lineCount
	char := l.peek()
	col := l.prevCol

	for char != 0 && char != '\n' {
		val += string(char)

		char = l.advance()
	}

	if val == "" {
		return nil
	}

	return &Token{kind: MultiString, value: val, row: row, col: col}
}

func (l *Lexer) lexUrl() *Token {
	val := ""
	row := l.lineCount
	next := l.peekNext()
	char := l.peek()
	col := l.prevCol

	for char != 0 && !unicode.IsSpace(char) && char != '\n' && !(char == ')' && (next == '[' || next == ' ')) {
		val += string(char)

		char = l.advance()
		next = l.peekNext()
	}

	if val == "" {
		return nil
	}

	return &Token{kind: NarrowString, value: val, row: row, col: col}
}

func (l *Lexer) lexImgAlt() *Token {
	val := ""
	row := l.lineCount
	next := l.peekNext()
	char := l.peek()
	col := l.prevCol

	for char != 0 && char != '\n' && !(char == ']' && unicode.IsSpace(l.peekNext())) && !(char == ']' && (unicode.IsSpace(next) || next == 0)) {
		val += string(char)

		char = l.advance()
		next = l.peekNext()
	}

	if val == "" {
		return nil
	}

	return &Token{kind: MultiString, value: val, row: row, col: col}
}

func (l *Lexer) Lex() ([]*Token, error) {
	tokens := []*Token{}

	char := l.advance()
	for char != 0 {
		l.eatWhiteSpace()
		char = l.peek()
		startCol := l.prevCol

		switch char {
		case '#':
			l.mode = MultiMode
			if l.peekNext() == '#' {
				l.advance()

				if l.peekNext() == '#' {
					l.advance()

					tokens = append(
						tokens,
						&Token{kind: MetaHash, value: "###", row: l.lineCount, col: startCol},
					)
					l.mode = NarrowMode
				} else {
					tokens = append(tokens, &Token{kind: HeaderHash, value: "##", row: l.lineCount, col: startCol})
				}
			} else if l.peekNext() == '!' {
				l.advance()

				tokens = append(tokens, &Token{
					kind: HashBang, value: "#!", row: l.lineCount, col: startCol,
				})
			} else {
				tokens = append(
					tokens,
					&Token{kind: TitleHash, value: "#", row: l.lineCount, col: startCol},
				)
			}

			char = l.advance()
		case '(':
			tokens = append(
				tokens,
				&Token{kind: LeftParen, value: "(", row: l.lineCount, col: startCol},
			)
			l.mode = UrlMode

			char = l.advance()
		case ')':
			tokens = append(
				tokens,
				&Token{kind: RightParen, value: ")", row: l.lineCount, col: startCol},
			)

			char = l.advance()
		case '[':
			tokens = append(
				tokens,
				&Token{kind: LeftBrace, value: "[", row: l.lineCount, col: startCol},
			)
			l.mode = ImageMultiMode

			char = l.advance()
		case ']':
			tokens = append(
				tokens,
				&Token{kind: RightBrace, value: "]", row: l.lineCount, col: startCol},
			)

			char = l.advance()
		case '\n':
			tokens = append(
				tokens,
				&Token{kind: NewLine, value: "\n", row: l.lineCount, col: startCol},
			)

			l.mode = MultiMode
			char = l.advance()
		default:
			switch l.mode {
			case MultiMode:
				toke := l.lexMultiString()
				if toke != nil {
					tokens = append(tokens, toke)
				}
			case NarrowMode:
				toke := l.lexNarrowString()
				if toke != nil {
					tokens = append(tokens, toke)
				}
			case UrlMode:
				toke := l.lexUrl()
				if toke != nil {
					tokens = append(tokens, toke)
				}
			case ImageMultiMode:
				toke := l.lexImgAlt()
				if toke != nil {
					tokens = append(tokens, toke)
				}
			}

			char = l.peek()
		}
	}

	tokens = append(tokens, &Token{kind: EOF, value: "EOF", row: l.lineCount, col: l.col})

	return tokens, nil
}
