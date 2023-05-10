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

type DeclarationTypeClassique struct {
	Identifiant Ast
	TType       Ast
	Ctx         parser.DeclarationTypeClassiqueContext
	Type        *ttype.TigerType
}

func (e *DeclarationTypeClassique) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	e.Identifiant.VisitSemControl(slt, L)
	e.TType.VisitSemControl(slt, L)

	return &e.Ctx
}

func (e *DeclarationTypeClassique) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *DeclarationTypeClassique) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("DeclarationTypeClassique")

	id := e.Identifiant.Draw(g)
	g.CreateEdge("Id", node, id)

	typeNode := e.TType.Draw(g)
	g.CreateEdge("Type", node, typeNode)

	return node
}

func (l *AstCreatorListener) DeclarationTypeClassiqueEnter(ctx parser.DeclarationTypeClassiqueContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) DeclarationTypeClassiqueExit(ctx parser.DeclarationTypeClassiqueContext) {
	// Get back the last element of the stack
	declType := &DeclarationTypeClassique{
		Ctx: ctx,
	}

	declType.TType = l.PopAst()
	declType.Identifiant = l.PopAst()

	// Verify that the type is not already defined
	if _, err := l.Slt.GetSymbol(declType.Identifiant.(*Identifiant).Id); err == nil {
		l.Logger.NewSemanticError(logger.ErrorIdIsAlreadyDefinedInScope, declType.Identifiant.(*Identifiant).Ctx, declType.Identifiant.(*Identifiant).Id)
	}

	// Verify if the type is already defined
	if _, err := l.Slt.GetSymbol(declType.TType.(*Identifiant).Id); err != nil {
		l.Logger.NewSemanticError(logger.ErrorTypeIsNotDefined, declType.TType.(*Identifiant).Ctx, declType.TType.(*Identifiant).Id)
	}

	t := &slt.Symbol{
		Name: declType.Identifiant.(*Identifiant).Id,
		Type: declType.TType.ReturnType(),
	}
	l.Slt.AddSymbol(declType.Identifiant.(*Identifiant).Id, t)

	// Add the new type to the node
	declType.Type = t.Type

	l.PushAst(declType)
}

func (e *DeclarationTypeClassique) EnterAsm(writer *asm.AssemblyWriter) {
	defer e.ExitAsm(writer)
}

func (e *DeclarationTypeClassique) ExitAsm(writer *asm.AssemblyWriter) {
	// Nothing to do
}