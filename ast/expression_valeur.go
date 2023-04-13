package ast

import "ogtiger/parser"

type ExpressionValeur struct {
	Identifiant Ast
	ExpressionValeur2 Ast
}

func (e ExpressionValeur) Display() string {
	return " expressionValeur"
}

func  (l *AstCreatorListener) ExpressionValeurEnter(ctx parser.IExpressionValeurContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func  (l *AstCreatorListener) ExpressionValeurExit(ctx parser.IExpressionValeurContext) {
	if ctx.GetChildCount() == 1 {
		return
	}

	expressionValeur := &ExpressionValeur{}

	// Get the identifiant
	identifiant := l.AstStack[len(l.AstStack)-1]

	// Get the expressionValeur2
	expressionValeur2 := l.AstStack[len(l.AstStack)-2]

	// Pop the two last element of the stack
	l.AstStack = l.AstStack[:len(l.AstStack)-2]

	expressionValeur.Identifiant = identifiant
	expressionValeur.ExpressionValeur2 = expressionValeur2

	// Push the new element on the stack
	l.AstStack = append(l.AstStack, expressionValeur)
}