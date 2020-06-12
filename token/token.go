package token

type Type string

type Token struct {
	Type Type
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	IDENT = "IDENT" // add, foobar, x, y, ...
	INT = "INT"

	// 演算子
	ASSIGN = "="
	PLUS = "+"

	// delimiter
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keyword
	FUNCTION = "FUNCTION"
	LET = "LET"
)
