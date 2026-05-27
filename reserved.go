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

// Reserved statements declare a range of field numbers or field names that cannot be used in a message.
type Reserved struct {
	Position      scanner.Position
	Comment       *Comment
	Ranges        []Range
	FieldNames    []string
	InlineComment *Comment
	Parent        Visitee
}

// inlineComment is part of commentInliner.
func (r *Reserved) inlineComment(c *Comment) { _ = "STUB: not implemented"; return }

// Accept dispatches the call to the visitor.
func (r *Reserved) Accept(v Visitor) { _ = "STUB: not implemented"; return }

func (r *Reserved) parse(p *Parser) error { _ = "STUB: not implemented"; return nil }

// first char that determined tok

// use unread here because it could be start of ranges

func (r *Reserved) parent(v Visitee) { _ = "STUB: not implemented"; return }
