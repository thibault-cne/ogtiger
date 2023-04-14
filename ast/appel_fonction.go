package ast

import (
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type AppelFonction struct {
	Identifiant Ast
	Args        []Ast
	Ctx         parser.AppelFonctionContext
	Type        ttype.TigerType
}

func (e *AppelFonction) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *AppelFonction) Display() string {
	return " appel"
}

func (e *AppelFonction) Draw(prefix string, g *cgraph.Graph) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) AppelFonctionEnter(ctx parser.AppelFonctionContext) {
	// Nothing to do
}

func (l *AstCreatorListener) AppelFonctionExit(ctx parser.AppelFonctionContext) {
	appelFonction := &AppelFonction{
		Ctx: ctx,
	}

	// Get the identifiant
	expr := l.PopAst()

	appelFonction.Identifiant = expr

	// Get the args
	for i := 2; i < ctx.GetChildCount()-1; i += 2 {
		// Pop the last element of the stack
		// Add it to the args
		appelFonction.Args = append(appelFonction.Args, l.PopAst())
	}

	// Push the new element on the stack
	l.PushAst(appelFonction)
}
