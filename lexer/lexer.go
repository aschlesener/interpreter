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
	t := token.Token{Literal: string(l.currentChar), Type: getTokenType(l.currentChar)}
	l.readChar() // advance values in lexer

	return t
}

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

func getTokenType(tokenLiteral byte) token.TokenType {
	switch tokenLiteral {
	case '=':
		return token.ASSIGN
	case ';':
		return token.SEMICOLON
	case '(':
		return token.LPAREN
	case ')':
		return token.RPAREN
	case ',':
		return token.COMMA
	case '+':
		return token.PLUS
	case '{':
		return token.LBRACE
	case '}':
		return token.RBRACE
	case 0:
		return token.EOF
	}

	return token.EOF
}
