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

type DeclarationValeur struct {
	Id    Ast
	VType Ast
	Expr  Ast
	Ctx   parser.IDeclarationValeurContext
	Type  *ttype.TigerType
}

func (e *DeclarationValeur) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	e.Id.VisitSemControl(slt, L)
	if e.VType != nil {
		e.VType.VisitSemControl(slt, L)
	}
	e.Expr.VisitSemControl(slt, L)

	return e.Ctx
}

func (e *DeclarationValeur) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *DeclarationValeur) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("DeclarationValeur")

	id := e.Id.Draw(g)
	g.CreateEdge("Id", node, id)

	if e.VType != nil {
		typeNode := e.VType.Draw(g)
		g.CreateEdge("Type", node, typeNode)
	}

	expr := e.Expr.Draw(g)
	g.CreateEdge("Expr", node, expr)

	return node
}

func (l *AstCreatorListener) DeclarationValeurEnter(ctx parser.IDeclarationValeurContext) {
	// Nothing to do
}

func (l *AstCreatorListener) DeclarationValeurExit(ctx parser.IDeclarationValeurContext) {
	declarationValeur := &DeclarationValeur{
		Ctx: ctx,
	}
	var isNil = false

	declarationValeur.Expr = l.PopAst()
	declarationValeur.Type = declarationValeur.Expr.ReturnType()

	if len(ctx.AllIdentifiant()) > 1 {
		declarationValeur.VType = l.PopAst()
		declarationValeur.Type = declarationValeur.VType.ReturnType()

		// Verify that the type exists
		if _, err := l.Slt.GetSymbol(declarationValeur.VType.(*Identifiant).Id); err != nil {
			l.Logger.NewSemanticError(logger.ErrorTypeIsNotDefined, ctx, declarationValeur.VType.(*Identifiant).Id)
		}

		// Verify that the expr is not nil
		if ctx.GetChild(5).(*parser.ExpressionContext).GetText() == "nil" {
			isNil = true
		}
	} else {
		if ctx.GetChild(3).(*parser.ExpressionContext).GetText() == "nil" {
			isNil = true
		}
	}

	declarationValeur.Id = l.PopAst()

	if _, err := l.Slt.GetSymbolInScoope(declarationValeur.Id.(*Identifiant).Id); err == nil {
		l.Logger.NewSemanticError(logger.ErrorIdIsAlreadyDefinedInScope, ctx, declarationValeur.Id.(*Identifiant).Id)
	}

	// If isNil check if the type is a record type
	if isNil && !(declarationValeur.Type.ID == ttype.Record || declarationValeur.Type.ID == ttype.AnyRecord) {
		l.Logger.NewSemanticError(logger.ErrorNilIsOnlyForRecordTypes, ctx)
	}

	l.Slt.CreateSymbol(declarationValeur.Id.(*Identifiant).Id, declarationValeur.Type)
	
	l.PushAst(declarationValeur)
}

func (e *DeclarationValeur) EnterAsm(writer *asm.AssemblyWriter) {
	defer e.ExitAsm(writer)

	writer.Comment(fmt.Sprintf("Declarations de %s", e.Id.(*Identifiant).Id), 1)
	e.Expr.EnterAsm(writer)
}

func (e *DeclarationValeur) ExitAsm(writer *asm.AssemblyWriter) {
	// TODO: handle arrays and struct
	writer.Stmfd(string(asm.StackPointer), []string{string(asm.R8)})
	
	
	writer.SkipLine()
}
