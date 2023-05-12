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

type Program struct {
	Expr Ast
	Ctx  parser.IProgramContext
	Type *ttype.TigerType
}

func (e *Program) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	e.Expr.VisitSemControl(slt, L)
	return e.Ctx
}

func (e *Program) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *Program) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("Program")

	expr := e.Expr.Draw(g)
	g.CreateEdge("Expr", node, expr)

	return node
}

func (l *AstCreatorListener) ProgramEnter(ctx parser.IProgramContext) {
	// l.AstStack = append(l.AstStack, &Program{})
}

func (l *AstCreatorListener) ProgramExit(ctx parser.IProgramContext) {
	prog := &Program{
		Ctx:  ctx,
		Type: ttype.NewTigerType(ttype.NoReturn),
	}

	prog.Expr = l.PopAst()

	// Push the new element on the stack
	l.PushAst(prog)
}

func (e *Program) EnterAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	defer e.ExitAsm(writer, slt)

	// TODO: Add the lib functions
	writer.Label("main")
	writer.Comment("Display allocation", 1)
	writer.Mov("r0", fmt.Sprintf("#%d", slt.MaxScope() * 4), asm.NI)
	writer.Bl("malloc", asm.NI)
	writer.Mov("r10", "r0", asm.NI)

	writer.SkipLine()
	switch e.Expr.(type) {
	case *Definition:
		label := fmt.Sprintf("blk_%d_%d", e.Expr.(*Definition).Slt.Region, e.Expr.(*Definition).Slt.Scope)
		writer.Bl(label, asm.NI)
	}

	e.Expr.EnterAsm(writer, slt)
}

func (e *Program) ExitAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	writer.B("_exit", asm.NI)

	writer.NewRegion()
	writer.Print()

	writer.NewRegion()
	writer.Exit(0)
}
