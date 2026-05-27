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
	"text/scanner"
)

// Option is a protoc compiler option
type Option struct {
	Position   scanner.Position
	Comment    *Comment
	Name       string
	Constant   Literal
	IsEmbedded bool
	// AggregatedConstants is DEPRECATED. These Literals are populated into Constant.OrderedMap
	AggregatedConstants []*NamedLiteral
	InlineComment       *Comment
	Parent              Visitee
}

// parse reads an Option body
// ( ident | //... | "(" fullIdent ")" ) { "." ident } "=" constant ";"
func (o *Option) parse(p *Parser) error { _ = "STUB: not implemented"; return nil }

// check for =

// parse value

// values of an option can have illegal escape sequences
// for the standard Go scanner used by this package.

// aggregate
// consume {

// non aggregate

// https://protobuf.dev/reference/protobuf/proto3-spec/#option
func (o *Option) parseOptionName(p *Parser) error { _ = "STUB: not implemented"; return nil }

// check for dot
// none

// consume dot

// check for closing parenthesis

// put it back

// inlineComment is part of commentInliner.
func (o *Option) inlineComment(c *Comment) { _ = "STUB: not implemented"; return }

// Accept dispatches the call to the visitor.
func (o *Option) Accept(v Visitor) {
	_ = "STUB: not implemented"

	// Doc is part of Documented
	return
}

func (o *Option) Doc() *Comment {
	_ = "STUB: not implemented"

	// parseAggregate reads options written using aggregate syntax.
	// tLEFTCURLY { has been consumed
	return nil
}

func (o *Option) parseAggregate(p *Parser) error { _ = "STUB: not implemented"; return nil }

// reconstruct the old, deprecated field

func (o *Option) parent(v Visitee) { _ = "STUB: not implemented"; return }
