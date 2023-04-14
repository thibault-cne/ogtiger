package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz/cgraph"
)

type OperationSi struct {
	Cond Ast
	Then Ast
	Else Ast
	Ctx  parser.IOperationSiContext
}

func (e *OperationSi) Display() string {
	return " si"
}

func (e *OperationSi) Draw(g *cgraph.Graph) *cgraph.Node {
	node, _ := g.CreateNode("OperationSi")
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
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationSiExit(ctx parser.IOperationSiContext) {
	OperationSi := &OperationSi{
		Ctx: ctx,
	}

	OperationSi.Cond = l.PopAst()
	OperationSi.Then = l.PopAst()

	if ctx.GetChildCount() == 6 {
		OperationSi.Else = l.PopAst()
	}

	l.PushAst(OperationSi)
}
