package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz"
)

type DeclarationFontion struct {
	Id   Ast
	Type Ast
	Args []Ast
	Expr Ast
	Ctx  parser.IDeclarationFonctionContext
}

func (e *DeclarationFontion) Display() string {
	return " declarationFontion"
}

func (e *DeclarationFontion) Draw(prefix string, g *graphviz.Graphviz) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) DeclarationFontionEnter(ctx parser.IDeclarationFonctionContext) {
	// Nothing to do
}

func (l *AstCreatorListener) DeclarationFontionExit(ctx parser.IDeclarationFonctionContext) {
	declarationFontion := &DeclarationFontion{Ctx: ctx}

	declarationFontion.Id = l.PopAst()

	// Pop all args
	for i := 0; i < len(ctx.AllDeclarationChamp()); i++ {
		declarationFontion.Args = append(declarationFontion.Args, l.PopAst())
	}

	// Pop the type if it exists
	if len(ctx.AllIdentifiant()) > 1 {
		declarationFontion.Type = l.PopAst()
	}

	declarationFontion.Expr = l.PopAst()

	l.PushAst(declarationFontion)
}
