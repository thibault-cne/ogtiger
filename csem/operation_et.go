package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type OperationEt struct {
	Left  Ast
	Right Ast
	Ctx   parser.IOperationEtContext
	Type  *ttype.TigerType
}

func (e *OperationEt) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *OperationEt) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("OperationEt")

	left := e.Left.Draw(g)
	g.CreateEdge("Left", node, left)
	right := e.Right.Draw(g)
	g.CreateEdge("Right", node, right)

	return node
}

func (l *SemControlListener) OperationEtEnter(ctx parser.IOperationEtContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *SemControlListener) OperationEtExit(ctx parser.IOperationEtContext) {
	if ctx.GetChildCount() == 1 {
		return
	}

	node := l.PopAst()

	// Get the other exprEt
	for i := 0; i < len(ctx.AllOperationComparaison())-1; i++ {
		node = &OperationEt{
			Ctx:   ctx,
			Left:  node,
			Right: l.PopAst(),
		}
	}

	l.PushAst(node)
}
