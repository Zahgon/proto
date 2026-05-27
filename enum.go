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

// Enum definition consists of a name and an enum body.
type Enum struct {
	Position scanner.Position
	Comment  *Comment
	Name     string
	Elements []Visitee
	Parent   Visitee
}

// Accept dispatches the call to the visitor.
func (e *Enum) Accept(v Visitor) {
	_ = "STUB: not implemented"

	// Doc is part of Documented
	return
}

func (e *Enum) Doc() *Comment {
	_ = "STUB: not implemented"

	// addElement is part of elementContainer
	return nil
}

func (e *Enum) addElement(v Visitee) { _ = "STUB: not implemented"; return }

// elements is part of elementContainer
func (e *Enum) elements() []Visitee {
	_ = "STUB: not implemented"

	// takeLastComment is part of elementContainer
	// removes and returns the last element of the list if it is a Comment.
	return nil
}

func (e *Enum) takeLastComment(expectedOnLine int) (last *Comment) {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enum) parse(p *Parser) error { _ = "STUB: not implemented"; return nil }

// not merged?

// parent is part of elementContainer
func (e *Enum) parent(p Visitee) {
	_ = "STUB: not implemented"

	// EnumField is part of the body of an Enum.
	return
}

type EnumField struct {
	Position scanner.Position
	Comment  *Comment
	Name     string
	Integer  int
	// ValueOption is deprecated, use Elements instead
	ValueOption   *Option
	Elements      []Visitee // such as Option and Comment
	InlineComment *Comment
	Parent        Visitee
}

// elements is part of elementContainer
func (f *EnumField) elements() []Visitee {
	_ = "STUB: not implemented"

	// takeLastComment is part of elementContainer
	// removes and returns the last element of the list if it is a Comment.
	return nil
}

func (f *EnumField) takeLastComment(expectedOnLine int) (last *Comment) {
	_ = "STUB: not implemented"
	return nil
}

// Accept dispatches the call to the visitor.
func (f *EnumField) Accept(v Visitor) { _ = "STUB: not implemented"; return }

// inlineComment is part of commentInliner.
func (f *EnumField) inlineComment(c *Comment) { _ = "STUB: not implemented"; return }

// Doc is part of Documented
func (f *EnumField) Doc() *Comment { _ = "STUB: not implemented"; return nil }

func (f *EnumField) parse(p *Parser) error { _ = "STUB: not implemented"; return nil }

// update deprecated field with the last option found

// put back this token for scanning inline comment

// addElement is part of elementContainer
func (f *EnumField) addElement(v Visitee) { _ = "STUB: not implemented"; return }

func (f *EnumField) parent(v Visitee) {
	_ = "STUB: not implemented"

	// IsDeprecated returns true if the option "deprecated" is set with value "true".
	return
}

func (f *EnumField) IsDeprecated() bool { _ = "STUB: not implemented"; return false }
