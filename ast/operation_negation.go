package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type OperationNegation struct {
	Expr Ast
	Ctx  parser.IOperationNegationContext
	Type *ttype.TigerType
}

func (e *OperationNegation) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *OperationNegation) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("Negation")

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
