package ast

import (
	"fmt"
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
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("ListAcces")

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

	// Reverse the list
	for i, j := 0, len(listAcces.AccesChamps)-1; i < j; i, j = i+1, j-1 {
		listAcces.AccesChamps[i], listAcces.AccesChamps[j] = listAcces.AccesChamps[j], listAcces.AccesChamps[i]
	}

	// Get the identifiant
	listAcces.Identifiant = l.PopAst()
}
