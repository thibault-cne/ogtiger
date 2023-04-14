package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type OperationSi struct {
	Cond Ast
	Then Ast
	Else Ast
	Ctx  parser.IOperationSiContext
	Type *ttype.TigerType
}

func (e *OperationSi) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *OperationSi) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("Si")

	cond := e.Cond.Draw(g)
	g.CreateEdge("Cond", node, cond)

	then := e.Then.Draw(g)
	g.CreateEdge("Then", node, then)

	if e.Else != nil {
		Else := e.Else.Draw(g)
		g.CreateEdge("Else", node, Else)
	}

	return node
}

func (l *AstCreatorListener) OperationSiEnter(ctx parser.IOperationSiContext) {
	// Creata a new region
	l.Slt.CreateRegion()
}

func (l *AstCreatorListener) OperationSiExit(ctx parser.IOperationSiContext) {
	OperationSi := &OperationSi{
		Ctx: ctx,
	}

	if ctx.GetChildCount() == 6 {
		OperationSi.Else = l.PopAst()
	}

	OperationSi.Then = l.PopAst()
	OperationSi.Cond = l.PopAst()

	// Leave the region
	l.Slt = l.Slt.Parent

	l.PushAst(OperationSi)
}
