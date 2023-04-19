package ast

import (
	"fmt"
	"ogtiger/logger"
	"ogtiger/parser"
	"ogtiger/slt"
	"ogtiger/ttype"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-graphviz/cgraph"
)

type OperationOu struct {
	Left  Ast
	Right Ast
	Ctx   parser.IOperationOuContext
	Type  *ttype.TigerType
}

func (e *OperationOu) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	e.Left.VisitSemControl(slt, L)
	return e.Ctx
}

func (e *OperationOu) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *OperationOu) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("OperationOu")

	left := e.Left.Draw(g)
	g.CreateEdge("Left", node, left)
	right := e.Right.Draw(g)
	g.CreateEdge("Right", node, right)

	return node
}

func (l *AstCreatorListener) OperationOuEnter(ctx parser.IOperationOuContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationOuExit(ctx parser.IOperationOuContext) {
	if ctx.GetChildCount() == 1 {
		return
	}

	node := l.PopAst()

	// Get the other exprEt
	for i := 0; i < len(ctx.AllOperationEt())-1; i++ {
		node = &OperationOu{
			Ctx:   ctx,
			Left:  node,
			Right: l.PopAst(),
			Type:  node.ReturnType(),
		}
	}

	l.PushAst(node)
}
