package ast

import (
	"ogtiger/parser"
)

type DeclarationTypeClassique struct {
	Identifiant Ast
	Type        Ast
	Ctx         parser.DeclarationTypeClassiqueContext
}

func (e *DeclarationTypeClassique) Display() string {
	return " alias"
}

func (l *AstCreatorListener) DeclarationTypeClassiqueEnter(ctx parser.DeclarationTypeClassiqueContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) DeclarationTypeClassiqueExit(ctx parser.DeclarationTypeClassiqueContext) {
	// Get back the last element of the stack
	opDeclType := &DeclarationTypeClassique{
		Ctx: ctx,
	}

	opDeclType.Identifiant = l.PopAst()
	opDeclType.Type = l.PopAst()

	l.PushAst(opDeclType)
}
