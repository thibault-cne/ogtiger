package ast

import (
	"ogtiger/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Term struct {
	Left Factor
	Right []TermRight
}

type TermRight struct {
	Op string
	Right Factor
}

func  (l *AstCreatorListener) TermEnter(ctx parser.ITermContext) {
	// l.AstStack = append(l.AstStack, &Term{})
}

func  (l *AstCreatorListener) TermExit(ctx parser.ITermContext) {
	// Get back the last element of the stack
	term := &Term{}

	// Pop the last element of the stack
	// l.AstStack = l.AstStack[:len(l.AstStack)-1]

	// Get the first term
	factor := l.AstStack[len(l.AstStack)-1]

	// Pop the last element of the stack
	l.AstStack = l.AstStack[:len(l.AstStack)-1]

	term.Left = factor.(Factor)

	// Get minus and plus and term number
	for i := 1; i < (ctx.GetChildCount() - 1) / 2; i++ {
		right := TermRight{}

		right.Op = ctx.GetChild(2 * i).(antlr.TerminalNode).GetText()
		right.Right = l.AstStack[len(l.AstStack)-1].(Factor)

		// Pop the last element of the stack
		l.AstStack = l.AstStack[:len(l.AstStack)-1]

		term.Right = append(term.Right, right)
	}

	// Push the new element on the stack
	l.AstStack = append(l.AstStack, term)
}