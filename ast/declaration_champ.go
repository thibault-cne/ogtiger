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

type DeclarationChamp struct {
	Left  Ast
	Right Ast
	Ctx   parser.IDeclarationChampContext
	Type  *ttype.TigerType
	ErrorCount int
}

func (e *DeclarationChamp) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	e.Left.VisitSemControl(slt, L)
	e.Right.VisitSemControl(slt, L)

	return e.Ctx
}

func (e *DeclarationChamp) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *DeclarationChamp) GetErrorCount() int {
	return e.ErrorCount
}

func (e *DeclarationChamp) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("DeclarationChamp")

	left := e.Left.Draw(g)
	g.CreateEdge("Left", node, left)

	right := e.Right.Draw(g)
	g.CreateEdge("Rigth", node, right)

	return node
}

func (l *AstCreatorListener) DeclarationChampEnter(ctx parser.IDeclarationChampContext) {
	// Nothing to do
}

func (l *AstCreatorListener) DeclarationChampExit(ctx parser.IDeclarationChampContext) {
	declarationChamp := &DeclarationChamp{
		Ctx: ctx,
	}

	declarationChamp.Right = l.PopAst()
	declarationChamp.Left = l.PopAst()

	l.PushAst(declarationChamp)
}
