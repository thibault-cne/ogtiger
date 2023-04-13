package ast

import "ogtiger/parser"

type ExpressionUnaire struct {
	Expr Ast
}

func (e ExpressionUnaire) Display() string {
	return " expressionUnaire"
}

func  (l *AstCreatorListener) ExpressionUnaireEnter(ctx parser.IExpressionUnaireContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func  (l *AstCreatorListener) ExpressionUnaireExit(ctx parser.IExpressionUnaireContext) {
	expressionUnaire := &ExpressionUnaire{}

	if ctx.GetChildCount() == 1 {
		return
	}

	// Get the first expr
	expr := l.AstStack[len(l.AstStack)-1]

	// Pop the last element of the stack
	l.AstStack = l.AstStack[:len(l.AstStack)-1]

	expressionUnaire.Expr = expr

	// Push the new element on the stack
	l.AstStack = append(l.AstStack, expressionUnaire)
}