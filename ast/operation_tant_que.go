package ast

import (
	"fmt"
	"ogtiger/logger"
	"ogtiger/parser"
	"ogtiger/slt"
	"ogtiger/ttype"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-graphviz/cgraph"
)

type OperationTantQue struct {
	Cond  Ast
	Block Ast
	Ctx   parser.IOperationTantqueContext
	Type  *ttype.TigerType
}

func (e *OperationTantQue) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	// TODO: Fill this
	return e.Ctx
}

func (e *OperationTantQue) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *OperationTantQue) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("OperationTantQue")

	cond := e.Cond.Draw(g)
	block := e.Block.Draw(g)

	g.CreateEdge("Cond", node, cond)
	g.CreateEdge("Block", node, block)

	return node
}

func (l *AstCreatorListener) OperationTantQueEnter(ctx parser.IOperationTantqueContext) {
	// Create a new region
	l.Slt = l.Slt.CreateRegion()
}

func (l *AstCreatorListener) OperationTantQueExit(ctx parser.IOperationTantqueContext) {
	oT := &OperationTantQue{
		Ctx:  ctx,
		Type: ttype.NewTigerType(ttype.NoReturn),
	}

	oT.Block = l.PopAst()
	oT.Cond = l.PopAst()

	// Leave the region
	l.Slt = l.Slt.Parent

	l.PushAst(oT)
}
