package ast

import (
	"ogtiger/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-graphviz"
)

type ListAcces struct {
	Identifiant Ast
	AccesChamps []Ast
	Ctx         parser.ListeAccesContext
}

func (l *ListAcces) Display() string {
	return " listAcces"
}

func (e *ListAcces) Draw(prefix string, g *graphviz.Graphviz) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) ListAccesEnter(ctx parser.ListeAccesContext) {
	// Nothing to do
}

func (l *AstCreatorListener) ListAccesExit(ctx parser.ListeAccesContext) {
	count, maxCount := 1, ctx.GetChildCount()

	listAcces := &ListAcces{Ctx: ctx}

	// Get the identifiant
	listAcces.Identifiant = l.PopAst()

	// Get the accesChamps
	for count < maxCount {
		text := ctx.GetChild(count).(*antlr.TerminalNodeImpl).GetText()

		if text == "." {
			count += 1
		} else {
			count += 2
		}

		listAcces.AccesChamps = append(listAcces.AccesChamps, l.PopAst())

		count += 1
	}
}
