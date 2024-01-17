package parser_test

import (
	"testing"

	"github.com/Kreonn/monkey/internal/ast"
	"github.com/Kreonn/monkey/internal/lexer"
	"github.com/Kreonn/monkey/internal/parser"
)

type expectedIdentifier struct {
	id string
}

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 69420;
`
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements should contain 3 statements, only got %d statements", len(program.Statements))
	}

	tests := []expectedIdentifier{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.id) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let', got '%s'", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("s is not a *ast.LetStatement, got '%T'", stmt)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s', got '%s'", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name not '%s', got '%s'", name, letStmt.Name)
		return false
	}

	return true
}
