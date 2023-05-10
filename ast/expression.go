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

type Expression struct {
	Left  Ast
	Right Ast
	Ctx   parser.IExpressionContext
	Type  *ttype.TigerType
	ErrorCount int
}

func (e *Expression) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	leftCtx := e.Left.VisitSemControl(slt, L)
	e.Right.VisitSemControl(slt, L)

	if !e.Left.ReturnType().Equals(e.Right.ReturnType()) {
		L.NewSemanticError(logger.ErrorWrongTypesExpression, leftCtx, e.Left.ReturnType(), e.Right.ReturnType())
	}

	return e.Ctx
}

func (e *Expression) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *Expression) GetErrorCount() int {
	return e.ErrorCount
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

func (e *Expression) EnterAsm(writer *asm.AssemblyWriter) {
	defer e.ExitAsm(writer)
}

func (e *Expression) ExitAsm(writer *asm.AssemblyWriter) {
	// Nothing to do
}
