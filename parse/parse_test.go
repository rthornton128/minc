package parse_test

import (
	"strings"
	"testing"
	"text/scanner"
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
		t.Error("expected no errors received ", p.ErrorCount)
	}
}

func TestBadParse(t *testing.T) {
	src := strings.NewReader("$")

	var p parse.Parser
	p.Init("badparse.mc", src)
	p.Error = func(s *scanner.Scanner, msg string) {}
	go p.Parse()
	time.Sleep(1500)

	if p.ErrorCount == 0 {
		t.Error("expected errors got none")
	}
}

func TestBadParseWithHander(t *testing.T) {
	src := strings.NewReader("void main) {}")

	var p parse.Parser
	p.Init("badparse.mc", src)
	go p.Parse()
	time.Sleep(1500)

	if p.ErrorCount != 1 {
		t.Error("expected one error got ", p.ErrorCount)
	}
}
