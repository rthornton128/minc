// scan.go
//
// Copyright (c) 2016 Rob Thornton
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package scan_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/rthornton128/minc/scan"
)

func TestScan(t *testing.T) {
	// Do a sanity check to ensure that all the tokens we expect in a
	// minimal C program will be reported accurately
	src := strings.NewReader("void main() {}")
	expect := []struct {
		Lit string
		Off int
		Tok scan.Token
	}{
		{"void", 0, scan.Ident},
		{"main", 5, scan.Ident},
		{"(", 9, scan.LParen},
		{")", 10, scan.RParen},
		{"{", 12, scan.LBrace},
		{"}", 13, scan.RBrace},
		{"", 14, scan.EOF},
	}

	var s scan.Scanner
	s.Init("valid_scan.mc", src)
	go s.Scan()

	var i int
	for item := range s.Items {
		if expect[i].Lit != item.Lit {
			t.Errorf("expected %s got %s", expect[i].Lit, item.Lit)
		}
		if expect[i].Off != item.Pos.Offset {
			t.Errorf("expected %d got %d", expect[i].Off, item.Pos.Offset)
		}
		if expect[i].Tok != item.Tok {
			t.Errorf("expected %s got %s", expect[i].Tok, item.Tok)
		}
		i++
	}
}

func TestInvalidScans(t *testing.T) {
	// Many, if not all, of these character will be valid tokens at a later
	// date. That said, they are not currently part of the spec for minimal
	// C. This test ensures they are reported as invalid characters.
	var s scan.Scanner
	s.Init("invalid_scan.mc", strings.NewReader("[3* /\t\t:"))
	go s.Scan()

	for item := range s.Items {
		if item.Tok != scan.Invalid && item.Tok != scan.EOF {
			t.Errorf("expected %s got %s", scan.Invalid, item.Tok)
		}
	}
}

func TestReadError(t *testing.T) {
	// There are several errors that can be reported internally by the scanner
	// but that's left for testing by the text/scanner package. In this
	// particular case, we are testing to ensure the errHandler is working so
	// we create an intentional error. The Scanner in text/scanner requires
	// that the NUL character not be present within a stream.
	var s scan.Scanner
	s.Init("read_error.mc", bytes.NewReader([]byte{0}))
	go s.Scan()

	err := <-s.Errors
	if err == nil {
		t.Error("expected error got nil")
	}
}
