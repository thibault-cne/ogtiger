package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz"
)

type DeclarationArrayType struct {
	Identifiant Ast
	Type        Ast
	Ctx         parser.DeclarationArrayTypeContext
}

func (e *DeclarationArrayType) Display() string {
	return " array_type"
}

func (e *DeclarationArrayType) Draw(prefix string, g *graphviz.Graphviz) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) DeclarationArrayTypeEnter(ctx parser.DeclarationArrayTypeContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) DeclarationArrayTypeExit(ctx parser.DeclarationArrayTypeContext) {
	// Get back the last element of the stack
	declArrType := &DeclarationArrayType{
		Ctx: ctx,
	}

	declArrType.Identifiant = l.PopAst()
	declArrType.Type = l.PopAst()

	l.PushAst(declArrType)
}
