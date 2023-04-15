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

type InstanciationType struct {
	Identifiant  Ast
	Identifiants []Ast
	Expressions  []Ast
	Ctx          parser.InstanciationTypeContext
	Type         *ttype.TigerType
}

func (e *InstanciationType) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	// TODO: Fill this
	return &e.Ctx
}

func (e *InstanciationType) ReturnType() *ttype.TigerType {
	return e.Type
}

func (i *InstanciationType) Display() string {
	return " instanciationType"
}

func (e *InstanciationType) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("InstanciationType")

	identifiant := e.Identifiant.Draw(g)
	g.CreateEdge("Identifiant", node, identifiant)

	for i := 0; i < len(e.Identifiants); i++ {
		identifiant := e.Identifiants[i].Draw(g)
		g.CreateEdge("Identifiant", node, identifiant)

		expression := e.Expressions[i].Draw(g)
		g.CreateEdge("Expression", node, expression)
	}

	return node
}

func (l *AstCreatorListener) InstanciationTypeEnter(ctx parser.InstanciationTypeContext) {
	// Nothing to do
}

func (l *AstCreatorListener) InstanciationTypeExit(ctx parser.InstanciationTypeContext) {
	instanciationType := &InstanciationType{Ctx: ctx}

	// Get the identifiants and the expressions
	for i := 0; i < len(ctx.AllExpression()); i++ {
		// Pop the last element of the stack
		// Add it to the expressions
		instanciationType.Expressions = append([]Ast{l.PopAst()}, instanciationType.Expressions...)

		// Pop the last element of the stack
		// Add it to the identifiants
		instanciationType.Identifiants = append([]Ast{l.PopAst()}, instanciationType.Identifiants...)
	}

	// Get the identifiant
	instanciationType.Identifiant = l.PopAst()

	// Set the type
	instanciationType.Type = instanciationType.Identifiant.ReturnType()

	// Push the new element on the stack
	l.PushAst(instanciationType)
}
