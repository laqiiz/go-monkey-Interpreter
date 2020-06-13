package ast

import (
	"github.com/laqiiz/go-monkey-Interpreter/token"
	"testing"
)

func TestString(t *testing.T) {
	pgm := Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{
					Type:    token.LET,
					Literal: "let",
				},
				Name: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "myVar",
					},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "anotherVar",
					},
					Value: "anotherVar",
				},
			},
		},
	}

	if pgm.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String wrong. got=%q", pgm.String())
	}

}
