package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-graphviz/cgraph"
)

type OperationComparaison struct {
	Op    string
	Left  Ast
	Right Ast
	Ctx   parser.IOperationComparaisonContext
	Type  *ttype.TigerType
}

func (e *OperationComparaison) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *OperationComparaison) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("OperationComparaison")

	left := e.Left.Draw(g)
	g.CreateEdge("Left", node, left)
	right := e.Right.Draw(g)
	g.CreateEdge("Right", node, right)

	return node
}

func (l *AstCreatorListener) OperationComparaisonEnter(ctx parser.IOperationComparaisonContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) OperationComparaisonExit(ctx parser.IOperationComparaisonContext) {
	if ctx.GetChildCount() == 1 {
		return
	}

	opCompar := &OperationComparaison{
		Ctx: ctx,
	}

	// Get left and right from stack
	right := l.PopAst()
	left := l.PopAst()

	opCompar.Left = left
	opCompar.Right = right

	// Get operator
	opCompar.Op = ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()

	l.PushAst(opCompar)
}
