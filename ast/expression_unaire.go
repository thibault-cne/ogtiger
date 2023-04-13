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
	// Pop the last element of the stack
	expr := l.PopAst()

	expressionUnaire.Expr = expr

	// Push the new element on the stack
	l.PushAst(expressionUnaire)
}