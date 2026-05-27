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

// Proto represents a .proto definition
type Proto struct {
	Filename string
	Elements []Visitee
}

// Accept dispatches the call to the visitor.
func (proto *Proto) Accept(v Visitor) {
	_ = "STUB: not implemented"
	// As Proto is not (yet) a Visitee, we enumerate its elements instead
	// v.VisitProto(proto)
	return
}

// addElement is part of elementContainer
func (proto *Proto) addElement(v Visitee) { _ = "STUB: not implemented"; return }

// elements is part of elementContainer
func (proto *Proto) elements() []Visitee { _ = "STUB: not implemented"; return nil }

// takeLastComment is part of elementContainer
// removes and returns the last element of the list if it is a Comment.
func (proto *Proto) takeLastComment(expectedOnLine int) (last *Comment) {
	_ = "STUB: not implemented"
	return nil
}

// parse parsers a complete .proto definition source.
func (proto *Proto) parse(p *Parser) error { _ = "STUB: not implemented"; return nil }

// not merged?

// BEGIN proto2

// END proto2

// continue

func (proto *Proto) parent(v Visitee) {
	_ = "STUB: not implemented"

	// elementContainer unifies types that have elements.
	return
}

type elementContainer interface {
	addElement(v Visitee)
	elements() []Visitee
	takeLastComment(expectedOnLine int) *Comment
}
