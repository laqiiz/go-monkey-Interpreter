package parser

import (
	"fmt"
	"github.com/laqiiz/go-monkey-Interpreter/ast"
	"github.com/laqiiz/go-monkey-Interpreter/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	in := `

let x = 5
let y = 10
let foobar = 838383
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
	t.Helper() // エラー部分が分かりにくかったので追加

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

func TestIntegerLiteralExpression(t *testing.T) {
	in := "5;"
	p := New(lexer.New(in))
	pgm := p.ParseProgram()
	checkParserErrors(t, p)

	if len(pgm.Statements) != 1 {
		t.Fatalf("pgm has not enough statements. got=%d", len(pgm.Statements))
	}
	stmt, ok := pgm.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("pgm.Statemtns[0] is not ast.ExpressionStatement. got=%T", pgm.Statements[0])
	}
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5", literal.TokenLiteral())
	}

}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		in       string
		ope      string
		integerV interface{}
	}{
		{"!5;", "!", int64(5)},
		{"-15", "-", int64(15)},
		{"!true;", "!", true},
		{"!false;", "!", false},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.in)
		p := New(l)
		pgm := p.ParseProgram()
		checkParserErrors(t, p)

		if len(pgm.Statements) != 1 {
			t.Fatalf("pgm.Statemtns does not contain %d statements. got=%d\n", 1, len(pgm.Statements))
		}

		stmt, ok := pgm.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("pgm.Statements[0] is not ast.ExpressionStatement. got=%T", pgm.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}

		if exp.Operator != tt.ope {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.ope, exp.Operator)
		}

		switch tt.integerV.(type) {
		case int64:
			if !testIntegerLiteral(t, exp.Right, tt.integerV) {
				return
			}
		case bool:
			if !testBooleanLiteral(t, exp.Right, tt.integerV.(bool)) {
				return
			}
		}

	}

}

func testIntegerLiteral(t *testing.T, il ast.Expression, v interface{}) bool {
	t.Helper() // helper

	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLIteral. got=%T", il)
		return false
	}

	if integ.Value != v {
		t.Errorf("integ.Value not %d. got=%d", v, integ.Value)
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", v) {
		t.Errorf("integ.TokenLIteral not %d. got=%s", v, integ.TokenLiteral())
		return false
	}
	return true
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		in     string
		leftV  int64
		ope    string
		rightV int64
	}{
		{"5 + 5", 5, "+", 5},
		{"5 - 5", 5, "-", 5},
		{"5 * 5", 5, "*", 5},
		{"5 / 5", 5, "/", 5},
		{"5 > 5", 5, ">", 5},
		{"5 < 5", 5, "<", 5},
		{"5 == 5", 5, "==", 5},
		{"5 != 5", 5, "!=", 5},
	}

	for _, tt := range infixTests {
		p := New(lexer.New(tt.in))
		pgm := p.ParseProgram()
		checkParserErrors(t, p)

		if len(pgm.Statements) != 1 {
			t.Fatalf("program.Statements does not countain %d statements. got=%d\n", 1, len(pgm.Statements))
		}
		stmt, ok := pgm.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("pgm.Statements[0] is not ast.ExpressionStatement. got=%T", pgm.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("exp is not ast.InfixExpression. got=%T", stmt.Expression)
		}

		if !testIntegerLiteral(t, exp.Left, tt.leftV) {
			return
		}

		if exp.Operator != tt.ope {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.ope, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.rightV) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"true",
			"true",
		},
		{"false",
			"false",
		},
		{"3 > 5 == false",
			"((3 > 5) == false)",
		}, {"3 < 5 == true",
			"((3 < 5) == true)",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{"(5 + 5) * 2 ",
			"((5 + 5) * 2)",
		},
		{"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{"-(5 + 5)",
			"(-(5 + 5))",
		},
		{"!(true == true)",
			"(!(true == true))",
		},
	}

	for _, tt := range tests {
		p := New(lexer.New(tt.in))
		pgm := p.ParseProgram()
		checkParserErrors(t, p)

		got := pgm.String()
		if got != tt.want {
			t.Errorf("want=%q, got=%q", tt.want, got)
		}
	}
}

func testIdentifier(t *testing.T, exp ast.Expression, v string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}

	if ident.Value != v {
		t.Errorf("ident.Value not %s. got=%s", v, ident.Value)
		return false
	}

	if ident.TokenLiteral() != v {
		t.Errorf("ident.TokenLiteral not %s. got=%s", v, ident.TokenLiteral())
	}

	return true
}

func testLiteralExpression(t *testing.T, exp ast.Expression, want interface{}) bool {
	switch v := want.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testIdentifier(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	}

	t.Errorf("type of exp not handled. got=%T", exp)

	return false
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, v bool) bool {
	bo, ok := exp.(*ast.Boolean)
	if !ok {
		t.Errorf("exp not *ast.Booealn. got=%T", exp)
		return false
	}

	if bo.Value != v {
		t.Errorf("bo.Value not %t. got==%t", v, bo.Value)
		return false
	}

	if bo.TokenLiteral() != fmt.Sprintf("%t", v) {
		t.Errorf("bo.TokenLiteral not %t. got=%s", v, bo.TokenLiteral())
		return false
	}
	return true

}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, ope string, right interface{}) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.InfixExpression. got=%T(%s)", exp, exp)
		return false
	}

	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}
	if opExp.Operator != ope {
		t.Errorf("exp.Operator is not '%s'. got=%q", ope, opExp.Operator)
		return false
	}
	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}

	if opExp.Operator != ope {
		t.Errorf("exp.Operator is not '%s'. got=%q", ope, opExp.Operator)
		return false
	}

	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}
	return true

}

func TestBoolExpression(t *testing.T) {
	//TODO 省略
	//	in := `
	//true;
	//false;
	//let boobar = true;
	//let barfoo = false;
	//`
}
