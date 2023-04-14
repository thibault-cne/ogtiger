package ast

import (
	"ogtiger/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-graphviz/cgraph"
)

type ListAcces struct {
	Identifiant Ast
	AccesChamps []Ast
	Ctx         parser.ListeAccesContext
}

func (l *ListAcces) Display() string {
	return " listAcces"
}

func (e *ListAcces) Draw(g *cgraph.Graph) *cgraph.Node {
	node, _ := g.CreateNode("ListAcces")

	id := e.Identifiant.Draw(g)

	edge, _ := g.CreateEdge("", node, id)
	edge.SetLabel("Identifiant")

	for _, accesChamp := range e.AccesChamps {
		accesChampNode := accesChamp.Draw(g)
		edge, _ = g.CreateEdge("", node, accesChampNode)
		edge.SetLabel("AccesChamp")
	}

	return node
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
