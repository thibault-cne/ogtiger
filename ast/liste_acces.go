package ast

import (
	"fmt"
	"ogtiger/logger"
	"ogtiger/parser"
	"ogtiger/slt"
	"ogtiger/ttype"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-graphviz/cgraph"
)

type ListAcces struct {
	Identifiant Ast
	AccesChamps []Ast
	Ctx         parser.ListeAccesContext
	Type        *ttype.TigerType
}

func (e *ListAcces) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	e.Identifiant.VisitSemControl(slt, L)
	for _, accesChamp := range e.AccesChamps {
		accesChamp.VisitSemControl(slt, L)
	}
	return &e.Ctx
}

func (e *ListAcces) ReturnType() *ttype.TigerType {
	return e.Type
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

	listAcces := &ListAcces{
		Ctx: ctx,
	}

	// Get the accesChamps
	for count < maxCount {
		text := ctx.GetChild(count).(*antlr.TerminalNodeImpl).GetText()

		if text == "." {
			count += 1
		} else {
			count += 2
		}

		// Prepare the accesChamp
		listAcces.AccesChamps = append([]Ast{l.PopAst()}, listAcces.AccesChamps...)

		count += 1
	}

	// Get the identifiant
	listAcces.Identifiant = l.PopAst()

	// Push the listAcces
	l.PushAst(listAcces)
}
