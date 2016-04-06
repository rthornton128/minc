package parse_test

import (
	"strings"
	"testing"
	"time"

	"github.com/rthornton128/minc/parse"
)

func TestParse(t *testing.T) {
	src := strings.NewReader("void main() {}")

	var p parse.Parser
	p.Init("parse.mc", src)
	go p.Parse()
	time.Sleep(1500)

	if p.ErrorCount != 0 {
		t.Error("expected no errors, received", p.ErrorCount)
	}
}
