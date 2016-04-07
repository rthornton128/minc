package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/scanner"

	"github.com/rthornton128/minc/cgen"
	"github.com/rthornton128/minc/parse"
)

var linker = "ld"
var lflags = ""
var assembler = "as"
var aflags = ""

func fatal(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

func main() {
	// command line flag parsing
	flag.StringVar(&linker, "linker", linker, "external linker to execute")
	flag.StringVar(&lflags, "lflags", lflags, "linker flags")
	flag.StringVar(&assembler, "asm", assembler, "exernal assembler to execute")
	flag.StringVar(&aflags, "aflags", aflags, "assembler flags")
	flag.BoolVar(&cgen.IsGas, "gnuas", cgen.IsGas, "emit intel_syntax directive")
	flag.Parse()

	// verify only one argument was supplied and test whether the extension is
	// valid
	if flag.NArg() != 1 {
		fatal("expected single argument containing file to parse")
	}

	if filepath.Ext(flag.Arg(0)) != ".mc" {
		fatal("unknown file type:", filepath.Ext(flag.Arg(0)))
	}

	// open file, parse it and generate assembly
	f, err := os.Open(flag.Arg(0))
	if err != nil {
		fatal(err)
	}
	defer f.Close()

	var p parse.Parser
	p.Init(flag.Arg(0), f)
	p.Error = func(_ *scanner.Scanner, msg string) {
		fmt.Fprintf(os.Stderr, msg)
	}

	name := filepath.Base(flag.Arg(0))
	name = name[:len(name)-len(filepath.Ext(name))]
	outFile, err := os.Create(name + ".s")

	prog := p.Parse()
	gen := cgen.CGen{outFile}
	gen.Generate(prog)
	outFile.Close()

	// call assembler and linker to produce executable binary
	execAssembler(assembler, aflags, name)
	execLinker(linker, lflags, name)
}

func execAssembler(asm, flags, name string) {
	if flags == "" {
		flags = fmt.Sprintf("-o %s.o %s.s", name, name)
	}
	execProg(asm, flags, name+".s")
}

func execProg(cmd, flags, rmFile string) {
	c := exec.Command(cmd, strings.Split(flags, " ")...)
	if msg, err := c.CombinedOutput(); err != nil {
		fmt.Fprintln(os.Stderr, string(msg))
		fatal(err)
	}
	os.Remove(rmFile)
}
