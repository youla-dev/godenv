// Package parser implements a parser for the .env files.
package parser

import (
	"fmt"

	"github.com/youla-dev/godenv/ast"
	"github.com/youla-dev/godenv/token"
)

// Scanner converts a sequence of characters into a sequence of tokens.
type Scanner interface {
	NextToken() token.Token
}

// Parser takes a Scanner and builds an abstract syntax tree.
type Parser struct {
	scanner Scanner
	token   token.Token
}

// New returns new Parser.
func New(scanner Scanner) *Parser {
	return &Parser{
		scanner: scanner,
		token:   scanner.NextToken(),
	}
}

// Parse parses the .env file and returns an ast.Statement.
func (p *Parser) Parse() (ast.Statement, error) {
	var statements []ast.Statement

	for p.token.Type != token.EOF {
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}

		statements = append(statements, stmt)
	}

	file := &ast.FileStatement{
		Statements: statements,
	}

	return file, nil
}

func (p *Parser) parseStatement() (ast.Statement, error) {
	p.skipBlankLine()

	switch p.token.Type {
	case token.Identifier:
		return p.parseAssignStatement()
	case token.Comment:
		return p.parseCommentStatement()
	default:
		return nil, fmt.Errorf("unexpected statement: %s(%q)", p.token.Type, p.token.Literal)
	}
}

func (p *Parser) parseCommentStatement() (ast.Statement, error) {
	comment := &ast.CommentStatement{
		Value: p.token.Literal,
	}

	p.nextToken()

	return comment, nil
}

func (p *Parser) parseAssignStatement() (ast.Statement, error) {
	name := p.token.Literal
	p.nextToken()

	switch p.token.Type {
	case token.NewLine, token.EOF:
		return p.parseNakedAssign(name)
	case token.Assign:
		p.nextToken()

		switch p.token.Type {
		case token.NewLine, token.EOF:
			return p.parseNakedAssign(name)
		case token.Value, token.RawValue:
			return p.parseCompleteAssign(name)
		}
	}

	return nil, fmt.Errorf("unexpected token: %s(%s)", p.token.Type, p.token.Literal)
}

func (p *Parser) parseNakedAssign(name string) (ast.Statement, error) {
	p.nextToken()
	return &ast.AssignStatement{Name: name}, nil
}

func (p *Parser) parseCompleteAssign(name string) (ast.Statement, error) {
	value := p.token.Literal
	p.nextToken()

	switch p.token.Type {
	case token.NewLine, token.EOF:
		p.nextToken()
		return &ast.AssignStatement{Name: name, Value: value}, nil
	default:
		return nil, fmt.Errorf("unexpected token: %s(%s)", p.token.Type, p.token.Literal)
	}
}

func (p *Parser) skipBlankLine() {
	for p.token.Type == token.NewLine || p.token.Type == token.Space {
		p.nextToken()
	}
}

func (p *Parser) nextToken() {
	p.token = p.scanner.NextToken()
}
