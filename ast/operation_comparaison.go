package ast

import (
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-graphviz/cgraph"
)

type OperationComparaison struct {
	Left  Ast
	Right []*OperationComparaisonFD
	Ctx   parser.IOperationComparaisonContext
	Type  ttype.TigerType
}

func (e *OperationComparaison) ReturnType() ttype.TigerType {
	return e.Type
}

type OperationComparaisonFD struct {
	Op    string
	Right Ast
}

func (e *OperationComparaison) Display() string {
	return " comparaison"
}

func (e *OperationComparaison) Draw(g *cgraph.Graph) *cgraph.Node {
	node, _ := g.CreateNode("OperationComparaison")
	left := e.Left.Draw(g)
	g.CreateEdge("Left", node, left)

	for _, right := range e.Right {
		rightNode := right.Right.Draw(g)
		g.CreateEdge("Right", node, rightNode)
	}

	return node
}

func (l *AstCreatorListener) OperationComparaisonEnter(ctx parser.IOperationComparaisonContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) OperationComparaisonExit(ctx parser.IOperationComparaisonContext) {
	// Get back the last element of the stack
	opCompar := &OperationComparaison{
		Ctx: ctx,
	}

	if ctx.GetChildCount() == 1 {
		return
	}

	opCompar.Left = l.PopAst()

	// Get minus and plus and term number
	for i := 0; i < (ctx.GetChildCount()-1)/2; i++ {
		right := &OperationComparaisonFD{}

		right.Op = ctx.GetChild(2*i + 1).(*antlr.TerminalNodeImpl).GetText()
		right.Right = l.PopAst()

		opCompar.Right = append(opCompar.Right, right)
	}

	l.PushAst(opCompar)
}
