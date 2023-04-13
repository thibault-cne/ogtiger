package ast

import (
	"ogtiger/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type ListAcces struct {
	Identifiant Ast
	AccesChamps []Ast
}

func (l ListAcces) Display() string {
	return " listAcces"
}

func (l *AstCreatorListener) ListAccesEnter(ctx parser.ListeAccesContext) {
	// Nothing to do
}

func (l *AstCreatorListener) ListAccesExit(ctx parser.ListeAccesContext) {
	count, maxCount := 1, ctx.GetChildCount()

	listAcces := &ListAcces{}

	// Get the identifiant
	listAcces.Identifiant = l.PopAst()

	// Get the accesChamps
	for count < maxCount {
		text := ctx.GetChild(count).(*antlr.TerminalNodeImpl).GetText()

		if (text == ".") {
			count += 1
		} else {
			count += 2
		}

		listAcces.AccesChamps = append(listAcces.AccesChamps, l.PopAst())

		count += 1
	}
}