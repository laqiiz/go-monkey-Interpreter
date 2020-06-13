package parser

import (
	"github.com/laqiiz/go-monkey-Interpreter/ast"
	"github.com/laqiiz/go-monkey-Interpreter/lexer"
	"github.com/laqiiz/go-monkey-Interpreter/token"
)

type Parser struct {
	l         *lexer.Lexer // 字句解析器インスタンスへのポインタ
	curToken  token.Token  // 現在のトークン
	peekToken token.Token  // 次のトークン(curTokenで不足時にはpeekTokenを用いる）
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	pgm :=&ast.Program{}
	pgm.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			pgm.Statements = append(pgm.Statements, stmt)
		}
		p.nextToken()
	}

	return pgm
}


func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{
		Token: p.curToken,
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) { // let IDENT = .. の形じゃなかったらnilを返す（Let Statementじゃない）
		return nil
	}

	//TODO セミコロンに遭遇するまで四季を読み飛ばしている
	if !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) curTokenIs(t token.Type) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	return false
}

