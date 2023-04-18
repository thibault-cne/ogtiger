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

type OperationEt struct {
	Left  Ast
	Right Ast
	Ctx   parser.IOperationEtContext
	Type  *ttype.TigerType
	ErrorCount int
}

func (e *OperationEt) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	e.Left.VisitSemControl(slt, L)
	e.Right.VisitSemControl(slt, L)
	return e.Ctx
}

func (e *OperationEt) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *OperationEt) GetErrorCount() int {
	return e.ErrorCount
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

func (l *AstCreatorListener) OperationEtEnter(ctx parser.IOperationEtContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationEtExit(ctx parser.IOperationEtContext) {
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
