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

type OperationNegation struct {
	Expr Ast
	Ctx  parser.IOperationNegationContext
	Type *ttype.TigerType
}

func (e *OperationNegation) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	exprCtx := e.Expr.VisitSemControl(slt, L)

	if !e.Expr.ReturnType().Equals(ttype.NewTigerType(ttype.Int)) {
		L.NewSemanticError("Expression of the negation is not an integer", exprCtx)
	}

	return e.Ctx
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
