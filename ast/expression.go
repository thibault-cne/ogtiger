package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz/cgraph"
)

type Expression struct {
	Left  Ast
	Right Ast
	Ctx   parser.IExpressionContext
}

func (e *Expression) Display() string {
	return " expression"
}

func (e *Expression) Draw(g *cgraph.Graph) *cgraph.Node {
	node, _ := g.CreateNode("Expression")
	left := e.Left.Draw(g)
	g.CreateEdge("Left", node, left)

	right := e.Right.Draw(g)
	g.CreateEdge("Right", node, right)

	return node
}

func (l *AstCreatorListener) ExprEnter(ctx parser.IExpressionContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) ExprExit(ctx parser.IExpressionContext) {
	expr := &Expression{
		Ctx: ctx,
	}

	if ctx.GetChildCount() == 1 {
		return
	}

	expr.Left = l.PopAst()
	expr.Right = l.PopAst()

	l.PushAst(expr)
}
