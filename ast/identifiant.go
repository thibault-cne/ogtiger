package ast

import (
	"fmt"
	"ogtiger/asm"
	"ogtiger/logger"
	"ogtiger/parser"
	"ogtiger/slt"
	"ogtiger/ttype"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-graphviz/cgraph"
)

type Identifiant struct {
	Id   string
	Ctx  parser.IIdentifiantContext
	Type *ttype.TigerType
}

func (e *Identifiant) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	return e.Ctx
}

func (e *Identifiant) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *Identifiant) Draw(g *cgraph.Graph) *cgraph.Node {
	id := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(id)
	node.SetLabel(e.Id)

	return node
}

func (l *AstCreatorListener) IdentifiantEnter(ctx parser.IIdentifiantContext) {
	// Nothing to do
}

func (l *AstCreatorListener) IdentifiantExit(ctx parser.IIdentifiantContext) {
	identifiant := &Identifiant{
		Id:  ctx.GetText(),
		Ctx: ctx,
	}

	t, _ := l.Slt.GetSymbolType(ctx.GetText())

	identifiant.Type = t

	l.PushAst(identifiant)
}

func (e *Identifiant) EnterAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	defer e.ExitAsm(writer, slt)

	if (e.Type.ID == ttype.Type || e.Type.ID == ttype.Function) {
		return
	}

	s, err := slt.GetSymbol(e.Id)

	if err != nil {
		panic(err)
	}

	writer.Comment(fmt.Sprintf("Use the static chain to get the value of %s", e.Id), 1)
	writer.Ldr("R0", "R10", asm.NI, -((slt.Scope - s.Offset) * 4))
	writer.Add("R0", "R0", fmt.Sprintf("#%d", e.Type.SizeInStack()), asm.NI)
	writer.Ldr("R8", "R0", asm.NI, 0)
}

func (e *Identifiant) ExitAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	if (e.Type.ID != ttype.Type && e.Type.ID != ttype.Function) {
		writer.SkipLine()
	}
}
