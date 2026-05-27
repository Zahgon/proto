// Copyright (c) 2017 Ernest Micklei
//
// MIT License
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package proto

import (
	"io"
	"text/scanner"
)

// Parser represents a parser.
type Parser struct {
	debug         bool
	scanner       *scanner.Scanner
	buf           *nextValues
	scannerErrors []error
}

// nextValues is to capture the result of next()
type nextValues struct {
	pos scanner.Position
	tok token
	lit string
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser { _ = "STUB: not implemented"; return nil }

// handleScanError is called from the underlying Scanner
func (p *Parser) handleScanError(s *scanner.Scanner, msg string) { _ = "STUB: not implemented"; return }

// ignoreIllegalEscapesWhile is called for scanning constants of an option.
// Such content can have a syntax that is not acceptable by the Go scanner.
// This temporary installs a handler that ignores only one type of error: illegal char escape
func (p *Parser) ignoreIllegalEscapesWhile(block func()) {
	_ = "STUB: not implemented"
	// during block call change error handler
	return
}

// this catches both "illegal char escape" <= go1.12 and "invalid char escape" go1.13
// too bad there is no constant for this in scanner pkg

// restore

// Parse parses a proto definition. May return a parse or scanner error.
func (p *Parser) Parse() (*Proto, error) { _ = "STUB: not implemented"; return nil, nil }

// see if it was a scanner error

// Filename is for reporting. Optional.
func (p *Parser) Filename(f string) { _ = "STUB: not implemented"; return }

const stringWithSingleQuote = "'"

// next returns the next token using the scanner or drain the buffer.
func (p *Parser) next() (pos scanner.Position, tok token, lit string) {
	_ = "STUB: not implemented"

	// consume buf
	return *new(scanner.Position), *new(token), ""
}

// single quote needs additional scanning

// pre: first single quote has been read
func (p *Parser) nextSingleQuotedString() (pos scanner.Position, tok token, lit string) {
	_ = "STUB: not implemented"
	// Save current scanner mode and temporarily disable comment scanning
	// to prevent // inside single quotes from being treated as comments
	return *new(scanner.Position), *new(token), ""
}

// string inside single quote

// empty single quoted string

// scan for partial tokens until actual closing single-quote(') token

// end quote expected

func (p *Parser) ignoreErrorsWhile(block func()) {
	_ = "STUB: not implemented"
	// during block call change error handler which ignores it all
	return
}

// restore

// nextPut sets the buffer
func (p *Parser) nextPut(pos scanner.Position, tok token, lit string) {
	_ = "STUB: not implemented"
	return
}

func (p *Parser) unexpected(found, expected string, obj interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) nextInteger() (i int, err error) { _ = "STUB: not implemented"; return 0, nil }

// hex decode

// nextIdentifier consumes tokens which may have one or more dot separators (namespaced idents).
func (p *Parser) nextIdentifier() (pos scanner.Position, tok token, lit string) {
	_ = "STUB: not implemented"
	return *new(scanner.Position), *new(token), ""
}

// leading dot allowed

func (p *Parser) nextMessageLiteralFieldName() (pos scanner.Position, tok token, lit string) {
	_ = "STUB: not implemented"
	return *new(scanner.Position), *new(token), ""
}

// consume right square

// nextTypeName implements the Packages and Name Resolution for finding the name of the type.
// Valid examples:
// .google.protobuf.Empty
// stream T must return tSTREAM
// optional int32 must return tOPTIONAL
// Bogus must return Bogus
func (p *Parser) nextTypeName() (pos scanner.Position, tok token, lit string) {
	_ = "STUB: not implemented"
	return *new(scanner.Position), *new(token), ""
}

// leading dot allowed

// type can be namespaced more

// consume dot

func (p *Parser) nextIdent(keywordStartAllowed bool) (pos scanner.Position, tok token, lit string) {
	_ = "STUB: not implemented"
	return *new(scanner.Position), *new(token), ""
}

// can be keyword

// proceed with keyword as first literal

// see if identifier is namespaced

// consume dot

func (p *Parser) peekNonWhitespace() rune { _ = "STUB: not implemented"; return 0 }

// consume it

// https://protobuf.dev/reference/protobuf/proto3-spec/
func (p *Parser) nextFullIdent(keywordStartAllowed bool) (pos scanner.Position, tok token, lit string) {
	_ = "STUB: not implemented"
	return *new(scanner.Position), *new(token), ""
}

// can be keyword

// proceed with keyword as first literal

// consume dot
