package ast

import (
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type OperationTantQue struct {
	Cond  Ast
	Block Ast
	Ctx   parser.IOperationTantqueContext
	Type  ttype.TigerType
}

func (e *OperationTantQue) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *OperationTantQue) Display() string {
	return " tantque"
}

func (e *OperationTantQue) Draw(g *cgraph.Graph) *cgraph.Node {
	node, _ := g.CreateNode("OperationTantQue")

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

	oT.Cond = l.PopAst()
	oT.Block = l.PopAst()

	l.PushAst(oT)
}
