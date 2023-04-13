// Code generated from ./grammar/tiger.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // tiger

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by tigerParser.
type tigerVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by tigerParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by tigerParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by tigerParser#operationOu.
	VisitOperationOu(ctx *OperationOuContext) interface{}

	// Visit a parse tree produced by tigerParser#operationEt.
	VisitOperationEt(ctx *OperationEtContext) interface{}

	// Visit a parse tree produced by tigerParser#operationComparaison.
	VisitOperationComparaison(ctx *OperationComparaisonContext) interface{}

	// Visit a parse tree produced by tigerParser#operationAddition.
	VisitOperationAddition(ctx *OperationAdditionContext) interface{}

	// Visit a parse tree produced by tigerParser#operationMultiplication.
	VisitOperationMultiplication(ctx *OperationMultiplicationContext) interface{}

	// Visit a parse tree produced by tigerParser#expressionUnaire.
	VisitExpressionUnaire(ctx *ExpressionUnaireContext) interface{}

	// Visit a parse tree produced by tigerParser#Sequence.
	VisitSequence(ctx *SequenceContext) interface{}

	// Visit a parse tree produced by tigerParser#Negation.
	VisitNegation(ctx *NegationContext) interface{}

	// Visit a parse tree produced by tigerParser#ExpressionIdentifiant.
	VisitExpressionIdentifiant(ctx *ExpressionIdentifiantContext) interface{}

	// Visit a parse tree produced by tigerParser#AppelFonction.
	VisitAppelFonction(ctx *AppelFonctionContext) interface{}

	// Visit a parse tree produced by tigerParser#ListeAcces.
	VisitListeAcces(ctx *ListeAccesContext) interface{}

	// Visit a parse tree produced by tigerParser#InstanciationType.
	VisitInstanciationType(ctx *InstanciationTypeContext) interface{}

	// Visit a parse tree produced by tigerParser#SiAlors.
	VisitSiAlors(ctx *SiAlorsContext) interface{}

	// Visit a parse tree produced by tigerParser#SiAlorsSinon.
	VisitSiAlorsSinon(ctx *SiAlorsSinonContext) interface{}

	// Visit a parse tree produced by tigerParser#TantQue.
	VisitTantQue(ctx *TantQueContext) interface{}

	// Visit a parse tree produced by tigerParser#Pour.
	VisitPour(ctx *PourContext) interface{}

	// Visit a parse tree produced by tigerParser#definition.
	VisitDefinition(ctx *DefinitionContext) interface{}

	// Visit a parse tree produced by tigerParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by tigerParser#declarationType.
	VisitDeclarationType(ctx *DeclarationTypeContext) interface{}

	// Visit a parse tree produced by tigerParser#DeclarationTypeClassique.
	VisitDeclarationTypeClassique(ctx *DeclarationTypeClassiqueContext) interface{}

	// Visit a parse tree produced by tigerParser#DeclarationArrayType.
	VisitDeclarationArrayType(ctx *DeclarationArrayTypeContext) interface{}

	// Visit a parse tree produced by tigerParser#DeclarationRecordType.
	VisitDeclarationRecordType(ctx *DeclarationRecordTypeContext) interface{}

	// Visit a parse tree produced by tigerParser#declarationChamp.
	VisitDeclarationChamp(ctx *DeclarationChampContext) interface{}

	// Visit a parse tree produced by tigerParser#declarationFonction.
	VisitDeclarationFonction(ctx *DeclarationFonctionContext) interface{}

	// Visit a parse tree produced by tigerParser#declarationValeur.
	VisitDeclarationValeur(ctx *DeclarationValeurContext) interface{}

	// Visit a parse tree produced by tigerParser#ChaineChr.
	VisitChaineChr(ctx *ChaineChrContext) interface{}

	// Visit a parse tree produced by tigerParser#Entier.
	VisitEntier(ctx *EntierContext) interface{}

	// Visit a parse tree produced by tigerParser#Nil.
	VisitNil(ctx *NilContext) interface{}

	// Visit a parse tree produced by tigerParser#Break.
	VisitBreak(ctx *BreakContext) interface{}

	// Visit a parse tree produced by tigerParser#identifiant.
	VisitIdentifiant(ctx *IdentifiantContext) interface{}
}
