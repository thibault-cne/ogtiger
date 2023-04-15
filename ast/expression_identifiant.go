package ast

import (
	"fmt"
	"ogtiger/logger"
	"ogtiger/parser"
	"ogtiger/slt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-graphviz/cgraph"
)

type ExpressionIdentifiant struct{}

func (e *ExpressionIdentifiant) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	// TODO: Fill this
	return nil
}

func (e *ExpressionIdentifiant) Display() string {
	return " expressionValeur"
}

func (e *ExpressionIdentifiant) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("ExpressionIdentifiant")

	return node
}

func (l *AstCreatorListener) ExpressionIdentifiantEnter(ctx parser.ExpressionIdentifiantContext) {
	// Nothing to do
}

func (l *AstCreatorListener) ExpressionIdentifiantExit(ctx parser.ExpressionIdentifiantContext) {
	// Nothing to do
}
