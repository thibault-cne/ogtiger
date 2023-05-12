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

type Integer struct {
	Valeur string
	Ctx    parser.EntierContext
	Type   *ttype.TigerType
}

func (e *Integer) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	return &e.Ctx
}

func (e *Integer) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *Integer) Draw(g *cgraph.Graph) *cgraph.Node {
	id := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(id)
	node.SetLabel(e.Valeur)

	return node
}

func (l *AstCreatorListener) IntegerEnter(ctx parser.EntierContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) IntegerExit(ctx parser.EntierContext) {
	// Get back the last element of the stack
	it := &Integer{
		Ctx:  ctx,
		Type: ttype.NewTigerType(ttype.Int),
	}

	it.Valeur = ctx.GetText()

	l.PushAst(it)
}

func (e *Integer) EnterAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	defer e.ExitAsm(writer, slt)
}

func (e *Integer) ExitAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	writer.Mov("r8", fmt.Sprintf("#%s", e.Valeur), asm.NI)
}
