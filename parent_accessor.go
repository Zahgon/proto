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

func getParent(child Visitee) Visitee { _ = "STUB: not implemented"; return *new(Visitee) }

type parentAccessor struct {
	parent Visitee
}

func (p *parentAccessor) VisitMessage(m *Message) { _ = "STUB: not implemented"; return }

func (p *parentAccessor) VisitService(v *Service) { _ = "STUB: not implemented"; return }

func (p *parentAccessor) VisitSyntax(s *Syntax) { _ = "STUB: not implemented"; return }

func (p *parentAccessor) VisitPackage(pkg *Package) { _ = "STUB: not implemented"; return }

func (p *parentAccessor) VisitOption(o *Option) { _ = "STUB: not implemented"; return }

func (p *parentAccessor) VisitImport(i *Import) { _ = "STUB: not implemented"; return }

func (p *parentAccessor) VisitNormalField(i *NormalField) { _ = "STUB: not implemented"; return }

func (p *parentAccessor) VisitEnumField(i *EnumField) { _ = "STUB: not implemented"; return }

func (p *parentAccessor) VisitEnum(e *Enum) { _ = "STUB: not implemented"; return }

func (p *parentAccessor) VisitComment(e *Comment) { _ = "STUB: not implemented"; return }
func (p *parentAccessor) VisitOneof(o *Oneof)     { _ = "STUB: not implemented"; return }

func (p *parentAccessor) VisitOneofField(o *OneOfField) { _ = "STUB: not implemented"; return }

func (p *parentAccessor) VisitReserved(rs *Reserved) { _ = "STUB: not implemented"; return }

func (p *parentAccessor) VisitRPC(rpc *RPC) { _ = "STUB: not implemented"; return }

func (p *parentAccessor) VisitMapField(f *MapField) { _ = "STUB: not implemented"; return }

func (p *parentAccessor) VisitGroup(g *Group) { _ = "STUB: not implemented"; return }

func (p *parentAccessor) VisitExtensions(e *Extensions) { _ = "STUB: not implemented"; return }

func (p *parentAccessor) VisitEdition(e *Edition) { _ = "STUB: not implemented"; return }

func (p *parentAccessor) VisitProto(*Proto) { _ = "STUB: not implemented"; return }
