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

import "text/scanner"

// Package specifies the namespace for all proto elements.
type Package struct {
	Position      scanner.Position
	Comment       *Comment
	Name          string
	InlineComment *Comment
	Parent        Visitee
}

// Doc is part of Documented
func (p *Package) Doc() *Comment { _ = "STUB: not implemented"; return nil }

func (p *Package) parse(pr *Parser) error { _ = "STUB: not implemented"; return nil }

// Accept dispatches the call to the visitor.
func (p *Package) Accept(v Visitor) {
	_ = "STUB: not implemented"

	// inlineComment is part of commentInliner.
	return
}

func (p *Package) inlineComment(c *Comment) { _ = "STUB: not implemented"; return }

func (p *Package) parent(v Visitee) { _ = "STUB: not implemented"; return }
