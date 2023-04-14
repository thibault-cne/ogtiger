package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz/cgraph"
)

type ExpressionIdentifiant struct{}

func (e *ExpressionIdentifiant) Display() string {
	return " expressionValeur"
}

func (e *ExpressionIdentifiant) Draw(g *cgraph.Graph) *cgraph.Node {
	node, _ := g.CreateNode("ExpressionIdentifiant")

	return node
}

func (l *AstCreatorListener) ExpressionIdentifiantEnter(ctx parser.ExpressionIdentifiantContext) {
	// Nothing to do
}

func (l *AstCreatorListener) ExpressionIdentifiantExit(ctx parser.ExpressionIdentifiantContext) {
	// Nothing to do
}
