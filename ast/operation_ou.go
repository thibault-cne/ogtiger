package ast

import (
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type OperationOu struct {
	Left  Ast
	Right []Ast
	Ctx   parser.IOperationOuContext
	Type  ttype.TigerType
}

func (e *OperationOu) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *OperationOu) Display() string {
	return " ou"
}

func (e *OperationOu) Draw(g *cgraph.Graph) *cgraph.Node {
	node, _ := g.CreateNode("OperationOu")
	left := e.Left.Draw(g)
	g.CreateEdge("Left", node, left)

	for i, right := range e.Right {
		rightNode := right.Draw(g)
		g.CreateEdge("Right"+string(i), node, rightNode)
	}

	return node
}

func (l *AstCreatorListener) OperationOuEnter(ctx parser.IOperationOuContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationOuExit(ctx parser.IOperationOuContext) {
	operationOu := &OperationOu{
		Ctx: ctx,
	}

	if ctx.GetChildCount() == 1 {
		return
	}

	operationOu.Left = l.PopAst()

	for i := 0; i < len(ctx.AllOperationEt())-1; i++ {
		operationOu.Right = append(operationOu.Right, l.PopAst())
	}

	l.PushAst(operationOu)
}
