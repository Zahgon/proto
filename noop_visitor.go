// Copyright (c) 2022 Ernest Micklei
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

var _ Visitor = NoopVisitor{}

// NoopVisitor is a no-operation visitor that can be used when creating your own visitor that is interested in only one or a few types.
// It implements the Visitor interface.
type NoopVisitor struct{}

// VisitMessage is part of Visitor interface
func (n NoopVisitor) VisitMessage(m *Message) {
	_ = "STUB: not implemented"

	// VisitService is part of Visitor interface
	return
}

func (n NoopVisitor) VisitService(v *Service) {
	_ = "STUB: not implemented"

	// VisitSyntax is part of Visitor interface
	return
}

func (n NoopVisitor) VisitSyntax(s *Syntax) {
	_ = "STUB: not implemented"

	// VisitSyntax is part of Visitor interface
	return
}

func (n NoopVisitor) VisitEdition(e *Edition) {
	_ = "STUB: not implemented"

	// VisitPackage is part of Visitor interface
	return
}

func (n NoopVisitor) VisitPackage(p *Package) {
	_ = "STUB: not implemented"

	// VisitOption is part of Visitor interface
	return
}

func (n NoopVisitor) VisitOption(o *Option) {
	_ = "STUB: not implemented"

	// VisitImport is part of Visitor interface
	return
}

func (n NoopVisitor) VisitImport(i *Import) {
	_ = "STUB: not implemented"

	// VisitNormalField is part of Visitor interface
	return
}

func (n NoopVisitor) VisitNormalField(i *NormalField) {
	_ = "STUB: not implemented"

	// VisitEnumField is part of Visitor interface
	return
}

func (n NoopVisitor) VisitEnumField(i *EnumField) {
	_ = "STUB: not implemented"

	// VisitEnum is part of Visitor interface
	return
}

func (n NoopVisitor) VisitEnum(e *Enum) {
	_ = "STUB: not implemented"

	// VisitComment is part of Visitor interface
	return
}

func (n NoopVisitor) VisitComment(e *Comment) {
	_ = "STUB: not implemented"

	// VisitOneof is part of Visitor interface
	return
}

func (n NoopVisitor) VisitOneof(o *Oneof) {
	_ = "STUB: not implemented"

	// VisitOneofField is part of Visitor interface
	return
}

func (n NoopVisitor) VisitOneofField(o *OneOfField) {
	_ = "STUB: not implemented"

	// VisitReserved is part of Visitor interface
	return
}

func (n NoopVisitor) VisitReserved(r *Reserved) {
	_ = "STUB: not implemented"

	// VisitRPC is part of Visitor interface
	return
}

func (n NoopVisitor) VisitRPC(r *RPC) {
	_ = "STUB: not implemented"

	// VisitMapField is part of Visitor interface
	return
}

func (n NoopVisitor) VisitMapField(f *MapField) {
	_ = "STUB: not implemented"

	// VisitGroup is part of Visitor interface
	return
}

func (n NoopVisitor) VisitGroup(g *Group) {
	_ = "STUB: not implemented"

	// VisitExtensions is part of Visitor interface
	return
}

func (n NoopVisitor) VisitExtensions(e *Extensions) { _ = "STUB: not implemented"; return }
