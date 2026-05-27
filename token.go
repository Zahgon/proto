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

// token represents a lexical token.
type token int

const (
	// Special tokens
	tILLEGAL token = iota
	tEOF
	tWS

	// Literals
	tIDENT

	// Misc characters
	tSEMICOLON   // ;
	tCOLON       // :
	tEQUALS      // =
	tQUOTE       // "
	tSINGLEQUOTE // '
	tLEFTPAREN   // (
	tRIGHTPAREN  // )
	tLEFTCURLY   // {
	tRIGHTCURLY  // }
	tLEFTSQUARE  // [
	tRIGHTSQUARE // ]
	tCOMMENT     // /
	tLESS        // <
	tGREATER     // >
	tCOMMA       // ,
	tDOT         // .

	// Keywords
	keywordsStart
	tEDITION
	tSYNTAX
	tSERVICE
	tRPC
	tRETURNS
	tMESSAGE
	tIMPORT
	tPACKAGE
	tOPTION
	tREPEATED
	tWEAK
	tPUBLIC

	// special fields
	tONEOF
	tMAP
	tRESERVED
	tENUM
	tSTREAM

	// numbers (pos or neg, float)
	tNUMBER

	// BEGIN proto2
	tOPTIONAL
	tGROUP
	tEXTENSIONS
	tEXTEND
	tREQUIRED
	// END proto2
	keywordsEnd
)

// typeTokens exists for future validation
const typeTokens = "double float int32 int64 uint32 uint64 sint32 sint64 fixed32 sfixed32 sfixed64 bool string bytes"

// isKeyword returns if tok is in the keywords range
func isKeyword(tok token) bool { _ = "STUB: not implemented"; return false }

// isWhitespace checks for space,tab and newline
func isWhitespace(r rune) bool { _ = "STUB: not implemented"; return false }

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool { _ = "STUB: not implemented"; return false }

// isString checks if the literal is quoted (single or double).
func isString(lit string) bool { _ = "STUB: not implemented"; return false }

func isComment(lit string) bool { _ = "STUB: not implemented"; return false }

func isNumber(lit string) bool { _ = "STUB: not implemented"; return false }

const doubleQuoteRune = rune('"')

// unQuote removes one matching leading and trailing single or double quote.
//
// https://github.com/emicklei/proto/issues/103
// cannot use strconv.Unquote as this unescapes quotes.
func unQuote(lit string) (string, rune) { _ = "STUB: not implemented"; return "", 0 }

func asToken(literal string) token {
	_ = "STUB: not implemented"

	// delimiters
	return *new(token)
}

// words

// special fields

// proto2

// special cases
