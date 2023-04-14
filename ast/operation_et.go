package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz/cgraph"
)

type OperationEt struct {
	Left  Ast
	Right []Ast
	Ctx   parser.IOperationEtContext
}

func (e *OperationEt) Display() string {
	return " et"
}

func (e *OperationEt) Draw(g *cgraph.Graph) *cgraph.Node {
	node, _ := g.CreateNode("OperationEt")
	left := e.Left.Draw(g)
	g.CreateEdge("Left", node, left)

	for i, right := range e.Right {
		rightNode := right.Draw(g)
		g.CreateEdge("Right"+string(i), node, rightNode)
	}

	return node
}

func (l *AstCreatorListener) OperationEtEnter(ctx parser.IOperationEtContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationEtExit(ctx parser.IOperationEtContext) {
	operationEt := &OperationEt{
		Ctx: ctx,
	}

	if ctx.GetChildCount() == 1 {
		return
	}

	operationEt.Left = l.PopAst()

	// Get the other exprEt
	for i := 0; i < len(ctx.AllOperationComparaison())-1; i++ {
		operationEt.Right = append(operationEt.Right, l.PopAst())
	}

	l.PushAst(operationEt)
}
