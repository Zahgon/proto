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

// Message consists of a message name and a message body.
type Message struct {
	Position scanner.Position
	Comment  *Comment
	Name     string
	IsExtend bool
	Elements []Visitee
	Parent   Visitee
}

func (m *Message) groupName() string { _ = "STUB: not implemented"; return "" }

// parse expects ident { messageBody
func (m *Message) parse(p *Parser) error { _ = "STUB: not implemented"; return nil }

// parseMessageBody parses elements after {. It consumes the closing }
func parseMessageBody(p *Parser, c elementContainer) error { _ = "STUB: not implemented"; return nil }

// not merged?

// BEGIN proto2

// look ahead

// not a group, will be tFIELD

// END proto2 only

// continue

// tFIELD

// Accept dispatches the call to the visitor.
func (m *Message) Accept(v Visitor) {
	_ = "STUB: not implemented"

	// addElement is part of elementContainer
	return
}

func (m *Message) addElement(v Visitee) { _ = "STUB: not implemented"; return }

// elements is part of elementContainer
func (m *Message) elements() []Visitee { _ = "STUB: not implemented"; return nil }

func (m *Message) takeLastComment(expectedOnLine int) (last *Comment) {
	_ = "STUB: not implemented"
	return nil
}

// Doc is part of Documented
func (m *Message) Doc() *Comment { _ = "STUB: not implemented"; return nil }

func (m *Message) parent(v Visitee) { _ = "STUB: not implemented"; return }
