package parser

import (
	"fmt"
	"github.com/laqiiz/go-monkey-Interpreter/ast"
	"github.com/laqiiz/go-monkey-Interpreter/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	in := `

let x = 5;
let y = 10;
let foobar = 838383;
`

	p := New(lexer.New(in))
	pgm := p.ParseProgram()
	checkParserErrors(t, p)

	if len(pgm.Statements) != 3 {
		t.Logf("pgm.Statements: %+v", pgm.String())
		t.Fatalf("program.Statements does not countain 3 statements got=%d", len(pgm.Statements))
	}

	tests := []struct {
		wantIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("tests[%d]", i), func(t *testing.T) {
			stmt := pgm.Statements[i]
			if !testLetStatement(t, stmt, tt.wantIdentifier) {
				return
			}
		})
	}

}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.errors
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLIteral() not '%s'. got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	in := `
return 5;
return 10;
return 993322;
`
	p := New(lexer.New(in))

	pgm := p.ParseProgram()
	checkParserErrors(t, p)

	if len(pgm.Statements) != 3 {
		t.Fatalf("program.Statements does not countain 3 statements got=%d", len(pgm.Statements))
	}

	for _, stmt := range pgm.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return' go %q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	in := "foobar;"

	p := New(lexer.New(in))

	pgm := p.ParseProgram()
	checkParserErrors(t, p)

	if len(pgm.Statements) != 1 {
		t.Fatalf("program.Statements does not enough statemetns. got=%d", len(pgm.Statements))
	}
	stmt, ok := pgm.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. go=%T", stmt.Expression)
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar", ident.TokenLiteral())
	}

}
