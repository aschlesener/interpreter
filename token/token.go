package token

const (
	ILLEGAL = "ILLEGAL" // a token or character we dont know about
	EOF     = "EOF"     // signifies our parser that it can stop

	// Identifiers + literals
	IDENT = "IDENT" // variable names -  add, foobar, x, y, etc.
	INT   = "INT"   // 123

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string // this is the token value
}
