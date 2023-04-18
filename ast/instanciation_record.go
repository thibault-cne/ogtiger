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

type InstanciationRecord struct {
	Identifiant  Ast
	Identifiants []Ast
	Expressions  []Ast
	Ctx          parser.InstanciationRecordContext
	Type         *ttype.TigerType
	ErrorCount int
}

func (e *InstanciationRecord) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	e.Identifiant.VisitSemControl(slt, L)
	for _, identifiant := range e.Identifiants {
		identifiant.VisitSemControl(slt, L)
	}
	for _, expression := range e.Expressions {
		expression.VisitSemControl(slt, L)
	}
	return &e.Ctx
}

func (e *InstanciationRecord) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *InstanciationRecord) GetErrorCount() int {
	return e.ErrorCount
}

func (e *InstanciationRecord) Draw(g *cgraph.Graph) *cgraph.Node {
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

func (l *AstCreatorListener) InstanciationRecordEnter(ctx parser.InstanciationRecordContext) {
	// Nothing to do
}

func (l *AstCreatorListener) InstanciationRecordExit(ctx parser.InstanciationRecordContext) {
	instanciationType := &InstanciationRecord{Ctx: ctx}

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
