package ast

import (
	"fmt"
	"ogtiger/parser"

	"github.com/goccy/go-graphviz/cgraph"
)

type ExpressionIdentifiant struct{}

func (e *ExpressionIdentifiant) Display() string {
	return " expressionValeur"
}

func (e *ExpressionIdentifiant) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("ExpressionIdentifiant")

	return node
}

func (l *SemControlListener) ExpressionIdentifiantEnter(ctx parser.ExpressionIdentifiantContext) {
	// Nothing to do
}

func (l *SemControlListener) ExpressionIdentifiantExit(ctx parser.ExpressionIdentifiantContext) {
	// Nothing to do
}
