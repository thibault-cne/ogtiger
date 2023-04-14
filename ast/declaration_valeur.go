package ast

import (
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type DeclarationValeur struct {
	Id    Ast
	VType Ast
	Expr  Ast
	Ctx   parser.IDeclarationValeurContext
	Type  ttype.TigerType
}

func (e *DeclarationValeur) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *DeclarationValeur) Display() string {
	return " declarationValeur"
}

func (e *DeclarationValeur) Draw(prefix string, g *cgraph.Graph) {
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
		declarationValeur.VType = l.PopAst()
	}

	declarationValeur.Expr = l.PopAst()

	l.PushAst(declarationValeur)
}
