package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz/cgraph"
)

type DeclarationChamp struct {
	Left  Ast
	Rigth Ast
	Ctx   parser.IDeclarationChampContext
}

func (e *DeclarationChamp) Display() string {
	return " declarationChamp"
}

func (e *DeclarationChamp) Draw(prefix string, g *cgraph.Graph) {
	// TODO: Draw the AST
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
