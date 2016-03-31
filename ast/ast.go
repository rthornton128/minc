package ast

import "text/scanner"

type (
	Function struct {
		Type      *Identifier
		Name      *Identifier
		ParamList *ParamList
		StmtBlock *StmtBlock
	}

	Identifier struct {
		Pos scanner.Position
		Lit string
	}

	ParamList struct {
		RParen scanner.Position
		LParen scanner.Position
	}

	Program struct {
		Function *Function
	}

	StmtBlock struct {
		LBrace scanner.Position
		RBrace scanner.Position
	}
)
