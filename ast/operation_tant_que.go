package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type OperationTantQue struct {
	Cond  Ast
	Block Ast
	Ctx   parser.IOperationTantqueContext
	Type  *ttype.TigerType
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
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationTantQueExit(ctx parser.IOperationTantqueContext) {
	oT := &OperationTantQue{
		Ctx: ctx,
	}

	oT.Block = l.PopAst()
	oT.Cond = l.PopAst()

	l.PushAst(oT)
}
