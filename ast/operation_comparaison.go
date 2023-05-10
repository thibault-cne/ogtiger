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

type OperationComparaison struct {
	Op    string
	Left  Ast
	Right Ast
	Ctx   parser.IOperationComparaisonContext
	Type  *ttype.TigerType
}

func (e *OperationComparaison) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	e.Left.VisitSemControl(slt, L)
	e.Right.VisitSemControl(slt, L)

	if !e.Left.ReturnType().Equals(e.Right.ReturnType()) {
		L.NewSemanticError(logger.ErrorWrongTypesInComparison, e.Ctx, e.Left.ReturnType(), e.Right.ReturnType())
	}

	return e.Ctx
}

func (e *OperationComparaison) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *OperationComparaison) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("OperationComparaison")

	left := e.Left.Draw(g)
	g.CreateEdge("Left", node, left)
	right := e.Right.Draw(g)
	g.CreateEdge("Right", node, right)

	return node
}

func (l *AstCreatorListener) OperationComparaisonEnter(ctx parser.IOperationComparaisonContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) OperationComparaisonExit(ctx parser.IOperationComparaisonContext) {
	if ctx.GetChildCount() == 1 {
		return
	}

	opCompar := &OperationComparaison{
		Ctx: ctx,
	}

	// Get left and right from stack
	right := l.PopAst()
	left := l.PopAst()

	opCompar.Left = left
	opCompar.Right = right
	opCompar.Type = left.ReturnType()

	// Get operator
	opCompar.Op = ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()

	l.PushAst(opCompar)
}

func (e *OperationComparaison) EnterAsm(writer *asm.AssemblyWriter) {
	defer e.ExitAsm(writer)
}

func (e *OperationComparaison) ExitAsm(writer *asm.AssemblyWriter) {
	// Nothing to do
}
