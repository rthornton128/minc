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

func (c *CGen) Function(f *ast.Function) {
	c.emit("%s:", f.Name)
	c.emit("mov eax, 0")
	c.emit("ret")
}

func (c *CGen) Generate(p *ast.Program) {
	c.Function(p.Function)
}
