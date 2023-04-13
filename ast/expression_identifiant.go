package ast

import "ogtiger/parser"

type ExpressionIdentifiant struct {}

func (e ExpressionIdentifiant) Display() string {
	return " expressionValeur"
}

func  (l *AstCreatorListener) ExpressionValeurEnter(ctx parser.ExpressionIdentifiantContext) {
	// Nothing to do
}

func  (l *AstCreatorListener) ExpressionValeurExit(ctx parser.ExpressionIdentifiantContext) {
	// Nothing to do
}