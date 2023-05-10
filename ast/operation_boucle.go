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

type OperationBoucle struct {
	Start    Ast
	StartVal Ast
	EndVal   Ast
	Block    Ast
	Slt      *slt.SymbolTable
	Ctx      parser.IOperationBoucleContext
	Type     *ttype.TigerType
}

func (e *OperationBoucle) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	e.Start.VisitSemControl(e.Slt, L)
	startValCtx := e.StartVal.VisitSemControl(e.Slt, L)
	endValCtx := e.EndVal.VisitSemControl(e.Slt, L)
	e.Block.VisitSemControl(e.Slt, L)

	if !e.StartVal.ReturnType().Equals(ttype.NewTigerType(ttype.Int)) {
		L.NewSemanticError("Start value of the loop is not an integer", startValCtx)
	}

	if !e.EndVal.ReturnType().Equals(ttype.NewTigerType(ttype.Int)) {
		L.NewSemanticError("End value of the loop is not an integer", endValCtx)
	}

	return e.Ctx
}

func (e *OperationBoucle) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *OperationBoucle) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("OperationBoucle")

	start := e.Start.Draw(g)
	g.CreateEdge("Start", node, start)

	startVal := e.StartVal.Draw(g)
	g.CreateEdge("StartVal", node, startVal)

	endVal := e.EndVal.Draw(g)
	g.CreateEdge("EndVal", node, endVal)

	block := e.Block.Draw(g)
	g.CreateEdge("Block", node, block)

	return node
}

func (l *AstCreatorListener) OperationBoucleEnter(ctx parser.IOperationBoucleContext) {
	// Create a new TDS
	l.Slt = l.Slt.CreateRegion()
}

func (l *AstCreatorListener) OperationBoucleExit(ctx parser.IOperationBoucleContext) {
	oB := &OperationBoucle{
		Ctx:  ctx,
		Type: ttype.NewTigerType(ttype.NoReturn),
	}

	oB.Block = l.PopAst()
	oB.EndVal = l.PopAst()
	oB.StartVal = l.PopAst()
	oB.Start = l.PopAst()

	// Add the new element to the TDS
	id := oB.Start.(*Identifiant)
	l.Slt.CreateSymbol(id.Id, ttype.NewTigerType(ttype.Int))

	// Pop the TDS
	l.Slt = l.Slt.Parent
	oB.Slt = l.Slt

	// Push the new element on the stack
	l.PushAst(oB)
}

func (e *OperationBoucle) EnterAsm(writer *asm.AssemblyWriter) {
	defer e.ExitAsm(writer)
}

func (e *OperationBoucle) ExitAsm(writer *asm.AssemblyWriter) {
	// Nothing to do
}
