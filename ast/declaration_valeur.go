package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz"
)

type DeclarationValeur struct {
	Id   Ast
	Type Ast
	Expr Ast
	Ctx  parser.IDeclarationValeurContext
}

func (e *DeclarationValeur) Display() string {
	return " declarationValeur"
}

func (e *DeclarationValeur) Draw(prefix string, g *graphviz.Graphviz) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) DeclarationValeurEnter(ctx parser.IDeclarationValeurContext) {
	// Nothing to do
}

func (l *AstCreatorListener) DeclarationValeurExit(ctx parser.IDeclarationValeurContext) {
	declarationValeur := &DeclarationValeur{
		Ctx: ctx,
	}

	declarationValeur.Id = l.PopAst()

	if len(ctx.AllIdentifiant()) > 1 {
		declarationValeur.Type = l.PopAst()
	}

	declarationValeur.Expr = l.PopAst()

	l.PushAst(declarationValeur)
}
