package ast

import (
	"fmt"
	"ogtiger/asm"
	"ogtiger/logger"
	"ogtiger/parser"
	"ogtiger/slt"
	"ogtiger/ttype"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-graphviz/cgraph"
)

type ExpressionUnaire struct {
	Expr Ast
	Ctx  parser.IExpressionUnaireContext
	Type *ttype.TigerType
}

func (e *ExpressionUnaire) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	e.Expr.VisitSemControl(slt, L)
	return e.Ctx
}

func (e *ExpressionUnaire) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *ExpressionUnaire) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("ExpressionUnaire")

	expr := e.Expr.Draw(g)
	g.CreateEdge("Expr", node, expr)

	return node
}

func (l *AstCreatorListener) ExpressionUnaireEnter(ctx parser.IExpressionUnaireContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) ExpressionUnaireExit(ctx parser.IExpressionUnaireContext) {
	// Nothing to do
}

func (e *ExpressionUnaire) EnterAsm(writer *asm.AssemblyWriter) {
	defer e.ExitAsm(writer)
}

func (e *ExpressionUnaire) ExitAsm(writer *asm.AssemblyWriter) {
	// Nothing to do
}
