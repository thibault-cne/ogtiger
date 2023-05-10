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

type InstanciationRecord struct {
	Identifiant  Ast
	Identifiants []Ast
	Expressions  []Ast
	Ctx          parser.InstanciationRecordContext
	Type         *ttype.TigerType
}

func (e *InstanciationRecord) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	e.Identifiant.VisitSemControl(slt, L)
	for _, identifiant := range e.Identifiants {
		identifiant.VisitSemControl(slt, L)
	}
	for _, expression := range e.Expressions {
		expression.VisitSemControl(slt, L)
	}

	// Check types of the expressions
	for i := range e.Identifiants {
		fieldType, err := e.Type.GetRecordFieldType(e.Identifiants[i].(*Identifiant).Id)

		if err != nil {
			L.NewSemanticError(logger.ErrorFieldDoesNotExistInRecord, e.Identifiants[i].(*Identifiant).Ctx, e.Identifiants[i].(*Identifiant).Id, e.Identifiant.(*Identifiant).Id)
		} else if !fieldType.Equals(e.Expressions[i].ReturnType()) {
			L.NewSemanticError(logger.ErrorWrongTypeInInstanciationRecord, e.Identifiants[i].(*Identifiant).Ctx, fieldType, e.Expressions[i].ReturnType())
		}
	}

	return &e.Ctx
}

func (e *InstanciationRecord) ReturnType() *ttype.TigerType {
	return e.Type
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

func (e *InstanciationRecord) EnterAsm(writer *asm.AssemblyWriter) {
	defer e.ExitAsm(writer)
}

func (e *InstanciationRecord) ExitAsm(writer *asm.AssemblyWriter) {
	// Nothing to do
}
