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

// Service defines a set of RPC calls.
type Service struct {
	Position scanner.Position
	Comment  *Comment
	Name     string
	Elements []Visitee
	Parent   Visitee
}

// Accept dispatches the call to the visitor.
func (s *Service) Accept(v Visitor) {
	_ = "STUB: not implemented"

	// Doc is part of Documented
	return
}

func (s *Service) Doc() *Comment {
	_ = "STUB: not implemented"

	// addElement is part of elementContainer
	return nil
}

func (s *Service) addElement(v Visitee) { _ = "STUB: not implemented"; return }

// elements is part of elementContainer
func (s *Service) elements() []Visitee {
	_ = "STUB: not implemented"

	// takeLastComment is part of elementContainer
	// removes and returns the last elements of the list if it is a Comment.
	return nil
}

func (s *Service) takeLastComment(expectedOnLine int) (last *Comment) {
	_ = "STUB: not implemented"
	return nil
}

// parse continues after reading "service"
func (s *Service) parse(p *Parser) error { _ = "STUB: not implemented"; return nil }

// not merged?

func (s *Service) parent(v Visitee) {
	_ = "STUB: not implemented"

	// RPC represents an rpc entry in a message.
	return
}

type RPC struct {
	Position       scanner.Position
	Comment        *Comment
	Name           string
	RequestType    string
	StreamsRequest bool
	ReturnsType    string
	StreamsReturns bool
	Elements       []Visitee
	InlineComment  *Comment
	Parent         Visitee

	// Options field is DEPRECATED, use Elements instead.
	Options []*Option
}

// Accept dispatches the call to the visitor.
func (r *RPC) Accept(v Visitor) {
	_ = "STUB: not implemented"

	// Doc is part of Documented
	return
}

func (r *RPC) Doc() *Comment {
	_ = "STUB: not implemented"

	// inlineComment is part of commentInliner.
	return nil
}

func (r *RPC) inlineComment(c *Comment) { _ = "STUB: not implemented"; return }

// parse continues after reading "rpc"
func (r *RPC) parse(p *Parser) error { _ = "STUB: not implemented"; return nil }

// allow for inline comment parsing

// parse options

// not merged?

// addElement is part of elementContainer
func (r *RPC) addElement(v Visitee) { _ = "STUB: not implemented"; return }

// handle deprecated field

// elements is part of elementContainer
func (r *RPC) elements() []Visitee { _ = "STUB: not implemented"; return nil }

func (r *RPC) takeLastComment(expectedOnLine int) (last *Comment) {
	_ = "STUB: not implemented"
	return nil
}

func (r *RPC) parent(v Visitee) { _ = "STUB: not implemented"; return }
