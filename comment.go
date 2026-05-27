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

// Comment one or more comment text lines, either in c- or c++ style.
type Comment struct {
	Position scanner.Position
	// Lines are comment text lines without prefixes //, ///, /* or suffix */
	Lines      []string
	Cstyle     bool // refers to /* ... */,  C++ style is using //
	ExtraSlash bool // is true if the comment starts with 3 slashes
}

// newComment returns a comment.
func newComment(pos scanner.Position, lit string) *Comment { _ = "STUB: not implemented"; return nil }

type inlineComment struct {
	line       string
	extraSlash bool
}

// Accept dispatches the call to the visitor.
func (c *Comment) Accept(v Visitor) {
	_ = "STUB: not implemented"

	// Merge appends all lines from the argument comment.
	return
}

func (c *Comment) Merge(other *Comment) { _ = "STUB: not implemented"; return }

func (c Comment) hasTextOnLine(line int) bool { _ = "STUB: not implemented"; return false }

// Message returns the first line or empty if no lines.
func (c Comment) Message() string { _ = "STUB: not implemented"; return "" }

// commentInliner is for types that can have an inline comment.
type commentInliner interface {
	inlineComment(c *Comment)
}

// maybeScanInlineComment tries to scan comment on the current line ; if present then set it for the last element added.
func maybeScanInlineComment(p *Parser, c elementContainer) { _ = "STUB: not implemented"; return }

// see if there is an inline Comment

// seen comment and on same line and elements have been added

// if the last added element can have an inline comment then set it

// TODO skip multiline?

// takeLastCommentIfEndsOnLine removes and returns the last element of the list if it is a Comment
func takeLastCommentIfEndsOnLine(list []Visitee, line int) (*Comment, []Visitee) {
	_ = "STUB: not implemented"
	return nil, nil
}

// mergeOrReturnComment creates a new comment and tries to merge it with the last element (if is a comment and is on the next line).
func mergeOrReturnComment(elements []Visitee, lit string, pos scanner.Position) *Comment {
	_ = "STUB: not implemented"
	return nil
}

// last element must be a comment to merge

// do not merge c-style comments

// last comment has text on previous line
// TODO handle last line of file could be inline comment

// parent is part of elementContainer
func (c *Comment) parent(Visitee) {
	_ = "STUB: not implemented"

	// consumeCommentFor is for reading and taking all comment lines before the body of an element (starting at {)
	return
}

func consumeCommentFor(p *Parser, e elementContainer) { _ = "STUB: not implemented"; return }

// not merged?

// bit of recursion is fine
