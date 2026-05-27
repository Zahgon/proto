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

// Field is an abstract message field.
type Field struct {
	Position      scanner.Position
	Comment       *Comment
	Name          string
	Type          string
	Sequence      int
	Options       []*Option
	InlineComment *Comment
	Parent        Visitee
}

// inlineComment is part of commentInliner.
func (f *Field) inlineComment(c *Comment) { _ = "STUB: not implemented"; return }

// NormalField represents a field in a Message.
type NormalField struct {
	*Field
	Repeated bool
	Optional bool // proto2
	Required bool // proto2
}

func newNormalField() *NormalField { _ = "STUB: not implemented"; return nil }

// Accept dispatches the call to the visitor.
func (f *NormalField) Accept(v Visitor) { _ = "STUB: not implemented"; return }

// Doc is part of Documented
func (f *NormalField) Doc() *Comment {
	_ = "STUB: not implemented"

	// parse expects:
	// [ "repeated" | "optional" ] type fieldName "=" fieldNumber [ "[" fieldOptions "]" ] ";"
	return nil
}

func (f *NormalField) parse(p *Parser) error { _ = "STUB: not implemented"; return nil }

// proto2

// parseFieldAfterType expects:
// fieldName "=" fieldNumber [ "[" fieldOptions "]" ] ";
func parseFieldAfterType(f *Field, p *Parser, parent Visitee) error {
	_ = "STUB: not implemented"
	return nil
}

// allow keyword as field name

// continue as identifier

// found expected token

// put it back so we can use the generic nextInteger

// see if there are options

// consume options

func consumeFieldComments(f *Field, p *Parser) { _ = "STUB: not implemented"; return }

// no longer a comment, put it back

// TODO copy paste
func consumeOptionComments(o *Option, p *Parser) { _ = "STUB: not implemented"; return }

// no longer a comment, put it back

// MapField represents a map entry in a message.
type MapField struct {
	*Field
	KeyType string
}

func newMapField() *MapField { _ = "STUB: not implemented"; return nil }

// Accept dispatches the call to the visitor.
func (f *MapField) Accept(v Visitor) {
	_ = "STUB: not implemented"

	// Doc is part of Documented
	return
}

func (f *MapField) Doc() *Comment {
	_ = "STUB: not implemented"

	// parse expects:
	// mapField = "map" "<" keyType "," type ">" mapName "=" fieldNumber [ "[" fieldOptions "]" ] ";"
	// keyType = "int32" | "int64" | "uint32" | "uint64" | "sint32" | "sint64" |
	//
	//	"fixed32" | "fixed64" | "sfixed32" | "sfixed64" | "bool" | "string"
	return nil
}

func (f *MapField) parse(p *Parser) error { _ = "STUB: not implemented"; return nil }

func (f *Field) parent(v Visitee) { _ = "STUB: not implemented"; return }

const optionNameDeprecated = "deprecated"

// IsDeprecated returns true if the option "deprecated" is set with value "true".
func (f *Field) IsDeprecated() bool { _ = "STUB: not implemented"; return false }
