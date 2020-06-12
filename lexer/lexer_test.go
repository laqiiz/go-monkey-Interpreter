package lexer

import (
	"fmt"
	"github.com/laqiiz/go-monkey-Interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		wantType    token.Type
		wantLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)

	for i, tt := range tests {
		t.Run(fmt.Sprintf("tests[%d]", i), func(t *testing.T) {

			tok := l.NextToken()
			if tok.Type != tt.wantType {
				t.Fatalf("tokentype wrong. want=%q, got=%q", tt.wantType, tok.Type)
			}

			if tok.Literal != tt.wantLiteral {
				t.Fatalf("literal wrong. want=%q, got=%q", tt.wantLiteral, tok.Literal)
			}
		})

	}
}
