// Copyright (c) 2018 Ernest Micklei
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

// Handler is a type of function that accepts a Visitee.
type Handler func(v Visitee)

// Walk recursively pays a visit to all Visitees of a Proto and calls each handler with it.
func Walk(proto *Proto, handlers ...Handler) { _ = "STUB: not implemented"; return }

func walk(container elementContainer, handlers ...Handler) { _ = "STUB: not implemented"; return }

// WithImport returns a Handler that will call the apply function when the Visitee is an Import.
func WithImport(apply func(*Import)) Handler { _ = "STUB: not implemented"; return *new(Handler) }

// WithMessage returns a Handler that will call the apply function when the Visitee is a Message.
func WithMessage(apply func(*Message)) Handler { _ = "STUB: not implemented"; return *new(Handler) }

// WithOption returns a Handler that will call the apply function when the Visitee is a Option.
func WithOption(apply func(*Option)) Handler { _ = "STUB: not implemented"; return *new(Handler) }

// WithEnum returns a Handler that will call the apply function when the Visitee is a Enum.
func WithEnum(apply func(*Enum)) Handler { _ = "STUB: not implemented"; return *new(Handler) }

// WithOneof returns a Handler that will call the apply function when the Visitee is a Oneof.
func WithOneof(apply func(*Oneof)) Handler { _ = "STUB: not implemented"; return *new(Handler) }

// WithService returns a Handler that will call the apply function when the Visitee is a Service.
func WithService(apply func(*Service)) Handler { _ = "STUB: not implemented"; return *new(Handler) }

// WithRPC returns a Handler that will call the apply function when the Visitee is a RPC.
func WithRPC(apply func(*RPC)) Handler { _ = "STUB: not implemented"; return *new(Handler) }

// WithPackage returns a Handler that will call the apply function when the Visitee is a Package.
func WithPackage(apply func(*Package)) Handler { _ = "STUB: not implemented"; return *new(Handler) }

// WithNormalField returns a Handler that will call the apply function when the Visitee is a NormalField.
func WithNormalField(apply func(*NormalField)) Handler {
	_ = "STUB: not implemented"
	return *new(Handler)
}
