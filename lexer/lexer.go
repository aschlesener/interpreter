package lexer

import "interpreter/token"

// Lexer contains the original input string and keeps track of current and next positions.
// Lexer currently only supports ASCII characters.
type Lexer struct {
	input           string // `=+(){},;`
	currentPosition int    // 0
	currentChar     byte   // TODO: Change to rune, to support unicode //=
	nextPosition    int    // 1
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // intialize values in Lexer
	return l
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhitespace()

	switch l.currentChar {
	case '=':
		if l.peekChar() == '=' {
			ch := l.currentChar
			l.readChar()
			t = token.Token{Type: token.EQ, Literal: string(ch) + string(l.currentChar)}
		} else {
			t = token.NewToken(token.ASSIGN, l.currentChar)
		}
	case '+':
		t = token.NewToken(token.PLUS, l.currentChar)
	case '-':
		t = token.NewToken(token.MINUS, l.currentChar)
	case '!':
		if l.peekChar() == '=' {
			ch := l.currentChar
			l.readChar()
			t = token.Token{Type: token.NEQ, Literal: string(ch) + string(l.currentChar)}
		} else {
			t = token.NewToken(token.BANG, l.currentChar)
		}
	case '*':
		t = token.NewToken(token.ASTERISK, l.currentChar)
	case '/':
		t = token.NewToken(token.SLASH, l.currentChar)
	case '<':
		t = token.NewToken(token.LT, l.currentChar)
	case '>':
		t = token.NewToken(token.GT, l.currentChar)
	case ';':
		t = token.NewToken(token.SEMICOLON, l.currentChar)
	case '(':
		t = token.NewToken(token.LPAREN, l.currentChar)
	case ')':
		t = token.NewToken(token.RPAREN, l.currentChar)
	case ',':
		t = token.NewToken(token.COMMA, l.currentChar)

	case '{':
		t = token.NewToken(token.LBRACE, l.currentChar)
	case '}':
		t = token.NewToken(token.RBRACE, l.currentChar)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.currentChar) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdent(t.Literal)
			return t // return early since we've already called readChar to advance in readIdentifer
		} else if isDigit(l.currentChar) {
			t.Literal = l.readNumber()
			t.Type = token.INT
			return t // return early since we've already called readChar to advance in readIdentifer
		} else {
			t = token.NewToken(token.ILLEGAL, l.currentChar)
		}
	}

	l.readChar() // advance values in lexer

	return t
}

// readChar advances to the next character in input
func (l *Lexer) readChar() {
	// if we're at end of input, set character to ASCII code for "NUL" to indicate EOF
	if l.nextPosition >= len(l.input) {
		l.currentChar = 0
	} else {
		// read the next character
		l.currentChar = l.input[l.nextPosition]
	}

	l.currentPosition = l.nextPosition
	l.nextPosition += 1
}

// peekChar reads the next character in input without advancing position
func (l *Lexer) peekChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}

// read an identifier (string of letters) from the input
func (l *Lexer) readIdentifier() string {
	position := l.currentPosition
	for isLetter(l.currentChar) {
		l.readChar()
	}

	return l.input[position:l.currentPosition]
}

// read a number (string of digits) from the input
func (l *Lexer) readNumber() string {
	position := l.currentPosition
	for isDigit(l.currentChar) {
		l.readChar()
	}

	return l.input[position:l.currentPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.currentChar == ' ' || l.currentChar == '\t' || l.currentChar == '\n' || l.currentChar == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
