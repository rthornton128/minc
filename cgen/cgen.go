// Package cgen produces x86 instructions using Intel syntax. At current,
// it is intended to produce code for the GAS (GNU Assembler) and is likely
// not portable with other assemblers like: llvm assembler, nasm, masm, etc.
package cgen

import (
	"fmt"
	"io"
	"os"

	"github.com/rthornton128/minc/ast"
)

// CGen wraps an io.Writer to output the generated assembly to. If
// Writer is nil, it defaults to stdout
type CGen struct{ io.Writer }

func (c *CGen) emit(msg string, args ...interface{}) {
	fmt.Fprintf(c.Writer, msg, args...)
}

func (c *CGen) emitPreamble() {
	c.emit(".intel_syntax noprefix\n\n")
}

func (c *CGen) function(f *ast.Function) {
	c.emit("%s:\n", f.Name.Lit)
	c.emitExit(ExitSuccess) // always close with success
}

// Generate begins the code generation process.
func (c *CGen) Generate(p *ast.Program) {
	// ensure no panics ensue as a result of Writer being nil. output is sent
	// to stdout
	if c.Writer == nil {
		c.Writer = os.Stdout
	}
	c.emitPreamble()
	c.emit(".global %s\n\n", p.Function.Name.Lit)
	c.emit(".section .text\n")
	c.function(p.Function)
}
