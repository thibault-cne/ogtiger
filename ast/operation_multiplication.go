package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-graphviz/cgraph"
)

type OperationMultiplication struct {
	Left  Ast
	Right Ast
	Ctx   parser.IOperationMultiplicationContext
	Type  *ttype.TigerType
}

type OperationDivision struct {
	Left  Ast
	Right Ast
	Ctx   parser.IOperationMultiplicationContext
	Type  *ttype.TigerType
}

func (e *OperationMultiplication) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *OperationDivision) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *OperationMultiplication) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("*")

	left := e.Left.Draw(g)
	g.CreateEdge("Left", node, left)
	right := e.Right.Draw(g)
	g.CreateEdge("Right", node, right)

	return node
}

func (e *OperationDivision) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("/")

	left := e.Left.Draw(g)
	g.CreateEdge("Left", node, left)
	right := e.Right.Draw(g)
	g.CreateEdge("Right", node, right)

	return node
}

func (l *AstCreatorListener) OperationMultiplicationEnter(ctx parser.IOperationMultiplicationContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) OperationMultiplicationExit(ctx parser.IOperationMultiplicationContext) {
	if ctx.GetChildCount() == 1 {
		return
	}

	// Get back the elements needed from the stack
	elements := make([]Ast, 0)

	for i := 0; i < (ctx.GetChildCount()+1)/2; i++ {
		elements = append(elements, l.PopAst())
	}

	node := elements[len(elements)-1]
	elements = elements[:len(elements)-1]

	// TODO: Check if the type is correct

	// Get minus and plus and term number
	for i := 0; 2*i < (ctx.GetChildCount() - 1); i++ {
		switch ctx.GetChild(2*i + 1).(*antlr.TerminalNodeImpl).GetText() {
		case "*":
			temp := &OperationMultiplication{
				Ctx:   ctx,
				Left:  node,
				Right: elements[len(elements)-1],
				Type: ttype.NewTigerType(ttype.Int),
			}
			node = temp
		case "/":
			temp := &OperationDivision{
				Ctx:   ctx,
				Left:  node,
				Right: elements[len(elements)-1],
				Type: ttype.NewTigerType(ttype.Int),
			}
			node = temp
		}

		elements = elements[:len(elements)-1]
	}

	l.PushAst(node)
}
