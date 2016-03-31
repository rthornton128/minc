package parse_test

import (
	"strings"
	"testing"

	"github.com/rthornton128/minc/parse"
)

func TestParse(t *testing.T) {
	src := strings.NewReader("void main() {}")

	var p parse.Parser
	p.Init("parse.mc", src)
	go p.Parse()

	err := <-p.Errors
	if err != nil {
		t.Error(err)
	}
}
