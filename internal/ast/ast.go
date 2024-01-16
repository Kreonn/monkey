package ast

import "github.com/Kreonn/monkey/internal/token"

// Node represents a single Node of the AST
type Node interface {
	TokenLiteral() string
}

// Statement represents a Node that does not produce a value
type Statement interface {
	Node
	statementNode()
}

// Expression represents a Node that produces a value
type Expression interface {
	Node
	expressionNode()
}

// Program is the root Node of the AST
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement represents a variable assignation Statement
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier represents an user-inputted Expression
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
