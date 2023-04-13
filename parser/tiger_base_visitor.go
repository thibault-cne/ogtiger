// Code generated from ./grammar/tiger.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // tiger

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BasetigerVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasetigerVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitOperationOu(ctx *OperationOuContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitOperationEt(ctx *OperationEtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitOperationComparaison(ctx *OperationComparaisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitOperationAddition(ctx *OperationAdditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitOperationMultiplication(ctx *OperationMultiplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitExpressionUnaire(ctx *ExpressionUnaireContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitSequence(ctx *SequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitNegation(ctx *NegationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitExpressionIdentifiant(ctx *ExpressionIdentifiantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitAppelFonction(ctx *AppelFonctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitListeAcces(ctx *ListeAccesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitInstanciationType(ctx *InstanciationTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitSiAlors(ctx *SiAlorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitSiAlorsSinon(ctx *SiAlorsSinonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitTantQue(ctx *TantQueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitPour(ctx *PourContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitDefinition(ctx *DefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitDeclarationType(ctx *DeclarationTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitDeclarationTypeClassique(ctx *DeclarationTypeClassiqueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitDeclarationArrayType(ctx *DeclarationArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitDeclarationRecordType(ctx *DeclarationRecordTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitDeclarationChamp(ctx *DeclarationChampContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitDeclarationFonction(ctx *DeclarationFonctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitDeclarationValeur(ctx *DeclarationValeurContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitChaineChr(ctx *ChaineChrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitEntier(ctx *EntierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitNil(ctx *NilContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitBreak(ctx *BreakContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetigerVisitor) VisitIdentifiant(ctx *IdentifiantContext) interface{} {
	return v.VisitChildren(ctx)
}
