// scan.go
//
// Copyright (c) 2016 Rob Thornton
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

// Scan wraps itself around text/scanner to provide a more compiler friendly
// interface. The intent is to provide a smaller snapshot of the scanning
// process to make the learning process more approachable at an earlly stage
package scan

import (
	"fmt"
	"io"
	"text/scanner"
)

type Item struct {
	Lit string
	Pos scanner.Position
	Tok Token
}

// Scanner embeds text/scanner.Scanner and provides facilities to send
// scanized lexems via the Items channel. Any errors are sent via the
// Errors channel
type Scanner struct {
	scanner.Scanner
	Items  chan *Item
	Errors chan error
}

// emit sends a new item down the Items channel
func (s *Scanner) emit(t Token) {
	s.Items <- &Item{Lit: s.TokenText(), Pos: s.Position, Tok: t}
}

// errHandler reports the position and message of errors encountered by the
// scanner
func (s *Scanner) errHandler(_ *scanner.Scanner, msg string) {
	s.Errors <- fmt.Errorf("%s %s", s.Position, msg)
}

// Init initializes the embedded scanner and initializes the Items and
// Errors channels. Init must be called prior to using any other functions
// else undefined behaviour or panics may occur
func (s *Scanner) Init(fileName string, r io.Reader) {
	s.Items = make(chan *Item)
	s.Errors = make(chan error)

	s.Scanner.Init(r)
	s.Scanner.Error = s.errHandler
	s.Scanner.Filename = fileName
}

// Scan is intended to be run in a goroutine but this isn't strictly
// necessary. Scan will loop until EOF is returned, sending results or
// errors via Items and Errors channels respectively
func (s *Scanner) Scan() {
	var r rune
	for r != scanner.EOF {
		r = s.Scanner.Scan()
		switch r {
		case scanner.Ident:
			s.emit(Ident)
		case '{':
			s.emit(LBrace)
		case '}':
			s.emit(RBrace)
		case '(':
			s.emit(LParen)
		case ')':
			s.emit(RParen)
		case scanner.EOF:
			s.emit(EOF)
		default:
			s.emit(Invalid)
		}
	}
	close(s.Items)
	s.Errors <- nil
}
