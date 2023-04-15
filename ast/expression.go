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

type Expression struct {
	Left  Ast
	Right Ast
	Ctx   parser.IExpressionContext
	Type  *ttype.TigerType
}

func (e *Expression) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	// TODO: Fill this
	return e.Ctx
}

func (e *Expression) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *Expression) Display() string {
	return " expression"
}

func (e *Expression) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("Expression")

	left := e.Left.Draw(g)
	g.CreateEdge("Left", node, left)

	right := e.Right.Draw(g)
	g.CreateEdge("Right", node, right)

	return node
}

func (l *AstCreatorListener) ExprEnter(ctx parser.IExpressionContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) ExprExit(ctx parser.IExpressionContext) {
	if ctx.GetChildCount() == 1 {
		return
	}

	expr := &Expression{
		Ctx:  ctx,
		Type: ttype.NewTigerType(ttype.NoReturn),
	}

	expr.Right = l.PopAst()
	expr.Left = l.PopAst()

	l.PushAst(expr)
}
