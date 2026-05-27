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

// Oneof is a field alternate.
type Oneof struct {
	Position scanner.Position
	Comment  *Comment
	Name     string
	Elements []Visitee
	Parent   Visitee
}

// addElement is part of elementContainer
func (o *Oneof) addElement(v Visitee) { _ = "STUB: not implemented"; return }

// elements is part of elementContainer
func (o *Oneof) elements() []Visitee {
	_ = "STUB: not implemented"

	// takeLastComment is part of elementContainer
	// removes and returns the last element of the list if it is a Comment.
	return nil
}

func (o *Oneof) takeLastComment(expectedOnLine int) (last *Comment) {
	_ = "STUB: not implemented"
	return nil
}

// parse expects:
// oneofName "{" { oneofField | emptyStatement } "}"
func (o *Oneof) parse(p *Parser) error { _ = "STUB: not implemented"; return nil }

// not merged?

// TODO call takeLastComment instead?

// continue

// Accept dispatches the call to the visitor.
func (o *Oneof) Accept(v Visitor) {
	_ = "STUB: not implemented"

	// Doc is part of Documented
	return
}

func (o *Oneof) Doc() *Comment {
	_ = "STUB: not implemented"

	// OneOfField is part of Oneof.
	return nil
}

type OneOfField struct {
	*Field
}

func newOneOfField() *OneOfField { _ = "STUB: not implemented"; return nil }

// Accept dispatches the call to the visitor.
func (o *OneOfField) Accept(v Visitor) { _ = "STUB: not implemented"; return }

// Doc is part of Documented
// Note: although Doc() is defined on Field, it must be implemented here as well.
func (o *OneOfField) Doc() *Comment { _ = "STUB: not implemented"; return nil }

func (o *Oneof) parent(v Visitee) { _ = "STUB: not implemented"; return }
