package cgen_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/rthornton128/minc/cgen"
	"github.com/rthornton128/minc/parse"
)

func TestPreamble(t *testing.T) {
	src := strings.NewReader("void main() {}")

	var p parse.Parser
	p.Init("cgen.mc", src)
	go func() {
		for {
			<-p.Errors
		}
	}()
	prog := p.Parse()
	buf := new(bytes.Buffer)
	gen := cgen.CGen{buf}
	gen.Generate(prog)

	fmt.Println(buf)
}
