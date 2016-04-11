// scan.go
//
// Copyright (c) 2016 Rob Thornton
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package scan

// Token represents valid tokens the scanner may emit
type Token int

// List of valid token constants
const (
	Invalid Token = iota // Invalid/Unknown character
	EOF                  // End-of-File
	Func                 // Function
	Ident                // Identifier
	LBrace               // "{"
	RBrace               // "}"
	LParen               // "("
	RParen               // ")"
)

var tokens = map[Token]string{
	Invalid: "invalid",
	EOF:     "eof",
	Func:    "func",
	Ident:   "identifier",
	LBrace:  "{",
	RBrace:  "}",
	LParen:  "(",
	RParen:  ")",
}

// String returns a textual represented of a valid token
func (t Token) String() string { return tokens[t] }

func Tokenize(s string) Token {
	if s == "func" {
		return Func
	}
	return Ident
}
