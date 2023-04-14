package ast

import (
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type DeclarationFontion struct {
	Id    Ast
	FType Ast
	Args  []Ast
	Expr  Ast
	Ctx   parser.IDeclarationFonctionContext
	Type  ttype.TigerType
}

func (e *DeclarationFontion) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *DeclarationFontion) Display() string {
	return " declarationFontion"
}

func (e *DeclarationFontion) Draw(prefix string, g *cgraph.Graph) {
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
		declarationFontion.FType = l.PopAst()
	}

	declarationFontion.Expr = l.PopAst()

	l.PushAst(declarationFontion)
}
