// Package ast provides an Abstract Syntax Tree for the MinC language
package ast

import "text/scanner"

type (
	// Function represents the core abstraction of a program. As the language
	// specification describes, it consists of a return type, symobolic name,
	// parameter list and a statement block
	Function struct {
		Type      *Identifier
		Name      *Identifier
		ParamList *ParamList
		StmtBlock *StmtBlock
	}

	// Identifier may be either a type name, like "void", or the main
	// funciton identifier, "main"
	Identifier struct {
		Pos scanner.Position
		Lit string
	}

	// ParamList represents the "()" portion of the function declaration
	// Neither are used but the position of both parentheses are retained
	// Since further compilation of the language do not use these positions
	// but they are there for scaffolding of later work
	ParamList struct {
		RParen scanner.Position
		LParen scanner.Position
	}

	// Program is the root of the AST. It has only one field, the main
	// function, and nothing else. In the future, it will be much more complex,
	// holding things like other functions and global variable declations
	Program struct {
		Function *Function
	}

	// StmtBlock represents the "{}" portion of the function declaration
	// As with ParamList, neither are used in the rest of the compilation
	// process
	StmtBlock struct {
		LBrace scanner.Position
		RBrace scanner.Position
	}
)
