package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz/cgraph"
)

type OperationNegation struct {
	Expr Ast
	Ctx  parser.IOperationNegationContext
}

func (e *OperationNegation) Display() string {
	return " negation"
}

func (e *OperationNegation) Draw(g *cgraph.Graph) *cgraph.Node {
	node, _ := g.CreateNode("OperationNegation")
	expr := e.Expr.Draw(g)
	g.CreateEdge("Expr", node, expr)

	return node
}

func (l *AstCreatorListener) OperationNegationEnter(ctx parser.IOperationNegationContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationNegationExit(ctx parser.IOperationNegationContext) {
	operationNegation := &OperationNegation{
		Ctx: ctx,
	}

	operationNegation.Expr = l.PopAst()

	l.PushAst(operationNegation)
}
