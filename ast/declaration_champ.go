package ast

import "ogtiger/parser"

type DeclarationChamp struct {
	Left Ast
	Rigth Ast
	Ctx  parser.IDeclarationChampContext
}

func (e DeclarationChamp) Display() string {
	return " declarationChamp"
}

func (l *AstCreatorListener) DeclarationChampEnter(ctx parser.IDeclarationChampContext) {
	// Nothing to do
}

func (l *AstCreatorListener) DeclarationChampExit(ctx parser.IDeclarationChampContext) {
	declarationChamp := &DeclarationChamp{
		Ctx: ctx,
	}

	declarationChamp.Left = l.PopAst()
	declarationChamp.Rigth = l.PopAst()

	l.PushAst(declarationChamp)
}