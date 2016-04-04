package cgen

import (
	"fmt"
	"io"

	"github.com/rthornton128/minc/ast"
)

type CGen struct{ io.Writer }

func (c *CGen) emit(msg string, args ...interface{}) {
	fmt.Fprintf(c.Writer, msg, args...)
}

func (c *CGen) emitPreamble() {
	c.emit(".intel_syntax noprefix\n\n")
}

func (c *CGen) Function(f *ast.Function) {
	c.emit("%s:\n", f.Name.Lit)
	c.emitExit(ExitSuccess) // always close with success
}

func (c *CGen) Generate(p *ast.Program) {
	c.emitPreamble()
	c.emit(".global %s\n\n", p.Function.Name.Lit)
	c.emit(".section .text\n")
	c.Function(p.Function)
}
