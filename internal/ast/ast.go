// Package ast declares the types used to represent syntax trees for the .env file.
package ast

// Node represents AST-node of the syntax tree.
type Node interface{}

// Statement represents syntax tree node of .env file statement (like: assignment or comment).
type Statement interface {
	Node
	statementNode()
}

// FileStatement node represents .env file statement, that contains assignments and comments.
type FileStatement struct {
	Statements []Statement
}

// AssignStatement node represents a assignment statement.
type AssignStatement struct {
	Name  string
	Value string
}

// CommentStatement node represents a comment statement.
type CommentStatement struct {
	Value string
}

func (s *FileStatement) statementNode()    {}
func (s *AssignStatement) statementNode()  {}
func (s *CommentStatement) statementNode() {}
