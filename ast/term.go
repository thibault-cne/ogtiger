package ast

import (
	"ogtiger/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Term struct {
	Left Ast
	Right []*TermRight
}

type TermRight struct {
	Op string
	Right Ast
}

func (t Term) Display() string {
	return " term"
}

func  (l *AstCreatorListener) TermEnter(ctx parser.ITermContext) {
	// l.AstStack = append(l.AstStack, &Term{})
}

func  (l *AstCreatorListener) TermExit(ctx parser.ITermContext) {
	// Get back the last element of the stack
	term := &Term{}

	// Pop the last element of the stack
	// l.AstStack = l.AstStack[:len(l.AstStack)-1]

	if ctx.GetChildCount() == 1 {
		return
	}

	// Get the first term
	factor := l.AstStack[len(l.AstStack)-1]

	// Pop the last element of the stack
	l.AstStack = l.AstStack[:len(l.AstStack)-1]

	term.Left = factor

	// Get minus and plus and term number
	for i := 0; i < (ctx.GetChildCount() - 1) / 2; i++ {
		right := &TermRight{}

		right.Op = ctx.GetChild(2 * i + 1).(*antlr.TerminalNodeImpl).GetText()
		right.Right = l.AstStack[len(l.AstStack)-1]

		// Pop the last element of the stack
		l.AstStack = l.AstStack[:len(l.AstStack)-1]

		term.Right = append(term.Right, right)
	}

	// Push the new element on the stack
	l.AstStack = append(l.AstStack, term)
}