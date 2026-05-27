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

// Group represents a (proto2 only) group.
// https://developers.google.com/protocol-buffers/docs/reference/proto2-spec#group_field
type Group struct {
	Position scanner.Position
	Comment  *Comment
	Name     string
	Optional bool
	Repeated bool
	Required bool
	Sequence int
	Elements []Visitee
	Parent   Visitee
}

// Accept dispatches the call to the visitor.
func (g *Group) Accept(v Visitor) {
	_ = "STUB: not implemented"

	// addElement is part of elementContainer
	return
}

func (g *Group) addElement(v Visitee) { _ = "STUB: not implemented"; return }

// elements is part of elementContainer
func (g *Group) elements() []Visitee {
	_ = "STUB: not implemented"

	// Doc is part of Documented
	return nil
}

func (g *Group) Doc() *Comment {
	_ = "STUB: not implemented"

	// takeLastComment is part of elementContainer
	// removes and returns the last element of the list if it is a Comment.
	return nil
}

func (g *Group) takeLastComment(expectedOnLine int) (last *Comment) {
	_ = "STUB: not implemented"
	return nil
}

// parse expects:
// groupName "=" fieldNumber { messageBody }
func (g *Group) parse(p *Parser) error { _ = "STUB: not implemented"; return nil }

func (g *Group) parent(v Visitee) { _ = "STUB: not implemented"; return }
