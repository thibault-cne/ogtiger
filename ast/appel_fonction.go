package ast

import "ogtiger/parser"

type AppelFonction struct {
	Identifiant Ast
	Args []Ast
}

func (a AppelFonction) Display() string {
	return " appelFonction"
}

func  (l *AstCreatorListener) AppelFonctionEnter(ctx parser.AppelFonctionContext) {
	// Nothing to do
}

func  (l *AstCreatorListener) AppelFonctionExit(ctx parser.AppelFonctionContext) {
	appelFonction := &AppelFonction{}

	// Get the identifiant
	expr := l.PopAst()

	appelFonction.Identifiant = expr

	// Get the args
	for i := 2; i < ctx.GetChildCount() - 1; i += 2 {
		// Pop the last element of the stack
		// Add it to the args
		appelFonction.Args = append(appelFonction.Args, l.PopAst())
	}

	// Push the new element on the stack
	l.PushAst(appelFonction)
}