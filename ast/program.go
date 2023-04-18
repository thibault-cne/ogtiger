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

type Program struct {
	Expr Ast
	Ctx  parser.IProgramContext
	Type *ttype.TigerType
	ErrorCount int
}

func (e *Program) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	e.Expr.VisitSemControl(slt, L)
	return e.Ctx
}

func (e *Program) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *Program) GetErrorCount() int {
	return e.ErrorCount
}

func (e *Program) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("Program")

	expr := e.Expr.Draw(g)
	g.CreateEdge("Expr", node, expr)

	return node
}

func (l *AstCreatorListener) ProgramEnter(ctx parser.IProgramContext) {
	// l.AstStack = append(l.AstStack, &Program{})
}

func (l *AstCreatorListener) ProgramExit(ctx parser.IProgramContext) {
	prog := &Program{
		Ctx:  ctx,
		Type: ttype.NewTigerType(ttype.NoReturn),
	}

	prog.Expr = l.PopAst()

	// Push the new element on the stack
	l.PushAst(prog)
}
