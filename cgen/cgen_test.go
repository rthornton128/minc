package cgen_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/rthornton128/minc/cgen"
	"github.com/rthornton128/minc/parse"
)

func TestGenerateSanity(t *testing.T) {
	src := strings.NewReader("func main() {}")

	var p parse.Parser
	p.Init("cgen.mc", src)

	// compile program
	prog := p.Parse()
	buf1 := new(bytes.Buffer)
	gen := cgen.CGen{buf1}
	gen.Generate(prog)

	// compile program again
	buf2 := new(bytes.Buffer)
	gen = cgen.CGen{buf2}
	gen.Generate(prog)

	// confirm results are the same
	if bytes.Compare(buf1.Bytes(), buf2.Bytes()) != 0 {
		t.Fatal("output from buf1 differs from buf2")
	}
}

func TestGenerateAssemble(t *testing.T) {
	src := strings.NewReader("func main() {}")

	var p parse.Parser
	p.Init("cgen.mc", src)

	// compile program
	prog := p.Parse()
	buf := new(bytes.Buffer)
	gen := cgen.CGen{buf}
	gen.Generate(prog)

	if msg, err := execAssembler(buf); err != nil {
		if msg != nil {
			t.Log(string(msg))
		}
		t.Fatal(err)
	}
}

func execAssembler(buf *bytes.Buffer) ([]byte, error) {
	f, err := ioutil.TempFile("", "cgentest")
	if err != nil {
		return nil, err
	}
	defer os.Remove(f.Name())
	defer f.Close()

	cmd := exec.Command("as", "-o", f.Name())
	cmd.Stdin = buf
	return cmd.CombinedOutput()
}
