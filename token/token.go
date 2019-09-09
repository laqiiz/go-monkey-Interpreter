package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF" // end of file

	// 識別子 + literal
	IDENT = "IDENT" // add, foobar, x,y, ...
	INT   = "INT"

	// operator
	ASSIGN = "="
	PLUS   = "+"

	// delimiter
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keyword
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
