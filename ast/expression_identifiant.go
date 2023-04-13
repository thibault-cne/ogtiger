package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz"
)

type ExpressionIdentifiant struct{}

func (e *ExpressionIdentifiant) Display() string {
	return " expressionValeur"
}

func (e *ExpressionIdentifiant) Draw(prefix string, g *graphviz.Graphviz) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) ExpressionIdentifiantEnter(ctx parser.ExpressionIdentifiantContext) {
	// Nothing to do
}

func (l *AstCreatorListener) ExpressionIdentifiantExit(ctx parser.ExpressionIdentifiantContext) {
	// Nothing to do
}
