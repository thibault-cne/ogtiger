package ast

import (
	"ogtiger/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Expr struct {
	Left Term
	Right []ExprRight
}

type ExprRight struct {
	Op string
	Right Term
}

func  (l *AstCreatorListener) ExprEnter(ctx parser.IExprContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func  (l *AstCreatorListener) ExprExit(ctx parser.IExprContext) {
	// Get back the last element of the stack
	expr := &Expr{}

	// Pop the last element of the stack
	// l.AstStack = l.AstStack[:len(l.AstStack)-1]

	// Get the first term
	term := l.AstStack[len(l.AstStack)-1]

	// Pop the last element of the stack
	l.AstStack = l.AstStack[:len(l.AstStack)-1]

	expr.Left = term.(Term)

	// Get minus and plus and term number
	for i := 1; i < (ctx.GetChildCount() - 1) / 2; i++ {
		right := ExprRight{}

		right.Op = ctx.GetChild(2 * i).(antlr.TerminalNode).GetText()
		right.Right = l.AstStack[len(l.AstStack)-1].(Term)

		// Pop the last element of the stack
		l.AstStack = l.AstStack[:len(l.AstStack)-1]

		expr.Right = append(expr.Right, right)
	}

	// Push the new element on the stack
	l.AstStack = append(l.AstStack, expr)
}