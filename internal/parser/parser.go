package parser

import (
	"github.com/Kreonn/monkey/internal/ast"
	"github.com/Kreonn/monkey/internal/lexer"
	"github.com/Kreonn/monkey/internal/token"
)

// Parser is the module responsible for building the AST,
// based on the Tokens produced by the Lexer
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New creates a new Parser
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

// ParseProgram builds the AST starting from the Program root Node,
// based on the Tokens produced by the Lexer
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = make([]ast.Statement, 0)

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO parse expressions
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(tt token.TokenType) bool {
	return p.curToken.Type == tt
}

func (p *Parser) peekTokenIs(tt token.TokenType) bool {
	return p.peekToken.Type == tt
}

func (p *Parser) expectPeek(tt token.TokenType) bool {
	if p.peekTokenIs(tt) {
		p.nextToken()
		return true
	} else {
		return false
	}
}
