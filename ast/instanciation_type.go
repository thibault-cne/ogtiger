package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz/cgraph"
)

type InstanciationType struct {
	Identifiant  Ast
	Identifiants []Ast
	Expressions  []Ast
	Ctx          parser.InstanciationTypeContext
}

func (i *InstanciationType) Display() string {
	return " instanciationType"
}

func (e *InstanciationType) Draw(g *cgraph.Graph) *cgraph.Node {
	node, _ := g.CreateNode("InstanciationType")

	identifiant := e.Identifiant.Draw(g)
	g.CreateEdge("Identifiant", node, identifiant)

	for i := 0; i < len(e.Identifiants); i++ {
		identifiant := e.Identifiants[i].Draw(g)
		g.CreateEdge("Identifiant", node, identifiant)

		expression := e.Expressions[i].Draw(g)
		g.CreateEdge("Expression", node, expression)
	}

	return node
}

func (l *AstCreatorListener) InstanciationTypeEnter(ctx parser.InstanciationTypeContext) {
	// Nothing to do
}

func (l *AstCreatorListener) InstanciationTypeExit(ctx parser.InstanciationTypeContext) {
	instanciationType := &InstanciationType{Ctx: ctx}

	// Get the identifiant
	identifiant := l.PopAst()

	instanciationType.Identifiant = identifiant

	// Get the identifiants count
	identifiantsCount := len(ctx.AllIdentifiant())

	// Get the identifiants and the expressions
	for i := 0; i < identifiantsCount; i++ {
		// Pop the last element of the stack
		// Add it to the identifiants
		instanciationType.Identifiants = append(instanciationType.Identifiants, l.PopAst())

		// Pop the last element of the stack
		// Add it to the expressions
		instanciationType.Expressions = append(instanciationType.Expressions, l.PopAst())
	}

	// Push the new element on the stack
	l.PushAst(instanciationType)
}
