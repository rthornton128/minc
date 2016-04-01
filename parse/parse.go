package parse

import (
	"fmt"
	"text/scanner"

	"github.com/rthornton128/minc/ast"
	"github.com/rthornton128/minc/scan"
)

type Parser struct {
	scan.Scanner
	item *scan.Item
}

func (p *Parser) Parse() {
	// start the scanner
	go p.Scan()

	// begin parsing, starting with the top-most non-terminal: program
	p.next()
	p.program()

	// a program must conclude with EOF
	p.expect(scan.EOF)
	p.Errors <- nil
}

// send an error
func (p *Parser) error(msg string, args ...interface{}) {
	p.Errors <- fmt.Errorf(msg, args...)
}

// expect returns the position of the lexem on success or generates
// an error on failure. in either instance, the scanner is advanced to
// the next token
func (p *Parser) expect(t scan.Token) scanner.Position {
	defer p.next()
	if p.item.Tok != t {
		p.error("expected %s got %s", t, p.item.Tok)
	}
	return p.item.Pos
}

// next pulls the next available item from the Items channel
func (p *Parser) next() {
	p.item = <-p.Items
}

func (p *Parser) program() *ast.Program {
	// a program consists of one, single function
	return &ast.Program{Function: p.function()}
}

func (p *Parser) function() *ast.Function {
	// a function has a strict order. the type name of void must come first.
	// the identifier main must come second. these are then followed by the
	// parameter list and statement block respectively.
	// neither the type name nor function name are verified at this point
	// this is handled later in the semantic analysis stage (type checking
	// of void) and linking (main being undeclared)
	return &ast.Function{
		Type:      p.identifier(),
		Name:      p.identifier(),
		ParamList: p.paramList(),
		StmtBlock: p.stmtBlock(),
	}
}

func (p *Parser) identifier() *ast.Identifier {
	// for an identifier, I want to record both its literal value and its
	// location. Again, the location is very important for error reporting
	// when it comes to verifying it against a symbol table
	return &ast.Identifier{
		Lit: p.item.Lit,
		Pos: p.expect(scan.Ident),
	}
}

func (p *Parser) paramList() *ast.ParamList {
	// the parameter list, which is always empty, consists of an opening and
	// closing parenthesis. Their location is important for error reporting.
	return &ast.ParamList{
		LParen: p.expect(scan.LParen),
		RParen: p.expect(scan.RParen),
	}
}

func (p *Parser) stmtBlock() *ast.StmtBlock {
	// see the description for parameter list for more information
	return &ast.StmtBlock{
		LBrace: p.expect(scan.LBrace),
		RBrace: p.expect(scan.RBrace),
	}
}
