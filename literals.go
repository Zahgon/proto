// Copyright (c) 2025 Ernest Micklei
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

// Literal represents intLit,floatLit,strLit or boolLit or a nested structure thereof.
type Literal struct {
	Position scanner.Position
	Source   string
	IsString bool

	// It not nil then the entry is actually a comment with line(s)
	// modelled this way because Literal is not an elementContainer
	Comment *Comment

	// The rune use to delimit the string value (only valid iff IsString)
	QuoteRune rune

	// literal value can be an array literal value (even nested)
	Array []*Literal

	// literal value can be a map of literals (even nested)
	// DEPRECATED: use OrderedMap instead
	Map map[string]*Literal

	// literal value can be a map of literals (even nested)
	// this is done as pairs of name keys and literal values so the original ordering is preserved
	OrderedMap LiteralMap
}

var emptyRune rune

// LiteralMap is like a map of *Literal but preserved the ordering.
// Can be iterated yielding *NamedLiteral values.
type LiteralMap []*NamedLiteral

// Get returns a Literal from the map.
func (m LiteralMap) Get(key string) (*Literal, bool) { _ = "STUB: not implemented"; return nil, false }

// exit on the first match

// SourceRepresentation returns the source (use the same rune that was used to delimit the string).
func (l Literal) SourceRepresentation() string { _ = "STUB: not implemented"; return "" }

// parse expects to read a literal constant after =.
func (l *Literal) parse(p *Parser) error { _ = "STUB: not implemented"; return nil }

// handle special element inside literal, a comment line

// peek at next token to see if it's structural (indicating end of current context)
// if so, don't recurse, just return with comment set

// put it back and return - let the caller handle these structural tokens

// put it back and continue with remaining entries (could be another comment or a literal)

// collect array elements

// if it's an empty array, consume the close bracket, set the Array to
// an empty array, and return

// if this is a comment-only literal, don't add it to array
// but keep reading to attach comment to next literal

// check what comes next

// array ends with comment, just break

// comma after comment, continue to next element

// put back the token

// continue loop - next iteration will parse the real literal
// and the comment will already be in 'e' from the recursive parse() call

// handle comments inside arrays

// negative number

// modify source and position

// peek for multiline strings

// NamedLiteral associates a name with a Literal
type NamedLiteral struct {
	*Literal
	Name string
	// PrintsColon is true when the Name must be printed with a colon suffix
	PrintsColon bool
}

// flatten the maps of each literal, recursively
// this func exists for deprecated Option.AggregatedConstants.
func collectAggregatedConstants(m map[string]*Literal) (list []*NamedLiteral) {
	_ = "STUB: not implemented"
	return nil
}

// sort list by position of literal

type byPosition []*NamedLiteral

func (b byPosition) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (b byPosition) Len() int      { _ = "STUB: not implemented"; return 0 }
func (b byPosition) Swap(i, j int) { _ = "STUB: not implemented"; return }

func parseAggregateConstants(p *Parser, container interface{}) (list []*NamedLiteral, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if tRIGHTSQUARE == tok {
// 	p.nextPut(pos, tok, lit)
// 	// caller has checked for open square ; will consume rightsquare, rightcurly and semicolon
// 	return
// }

// just consume it

//return

// assign to last parsed literal
// TODO: see TestUseOfSemicolonsInAggregatedConstants

// workaround issue #59 TODO

// concatenate with previous constant

// expect colon, aggregate or plain literal

// consume it

// see if nested aggregate is started

// create the map

// no aggregate, put back token

// now we see plain literal
