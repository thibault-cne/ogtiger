package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type OperationBoucle struct {
	Start    Ast
	StartVal Ast
	EndVal   Ast
	Block    Ast
	Ctx      parser.IOperationBoucleContext
	Type     *ttype.TigerType
}

func (e *OperationBoucle) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *OperationBoucle) Display() string {
	return " pour"
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
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationBoucleExit(ctx parser.IOperationBoucleContext) {
	oB := &OperationBoucle{
		Ctx: ctx,
	}

	oB.Block = l.PopAst()
	oB.EndVal = l.PopAst()
	oB.StartVal = l.PopAst()
	oB.Start = l.PopAst()

	// Push the new element on the stack
	l.AstStack = append(l.AstStack, oB)
}
