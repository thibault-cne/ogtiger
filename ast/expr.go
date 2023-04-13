package ast

import (
	"fmt"
	"ogtiger/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Expr struct {
	Left *Term
	Right []*ExprRight
}

type ExprRight struct {
	Op string
	Right *Term
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
	term := l.AstStack[len(l.AstStack)-1].(*Term)

	// Pop the last element of the stack
	l.AstStack = l.AstStack[:len(l.AstStack)-1]

	expr.Left = term

	// Get minus and plus and term number
	fmt.Printf("%d\n\n", ctx.GetChildCount())
	for i := 0; i < (ctx.GetChildCount() - 1) / 2; i++ {
		right := &ExprRight{}

		right.Op = ctx.GetChild(2 * i + 1).(*antlr.TerminalNodeImpl).GetText()
		right.Right = l.AstStack[len(l.AstStack)-1].(*Term)

		// Pop the last element of the stack
		l.AstStack = l.AstStack[:len(l.AstStack)-1]

		expr.Right = append(expr.Right, right)
	}

	// Push the new element on the stack
	l.AstStack = append(l.AstStack, expr)
}