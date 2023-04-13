package ast

import (
	"ogtiger/parser"
)

type Break struct {
	Ctx parser.BreakContext
}

func (e *Break) Display() string {
	return " break"
}

func (e *Break) Draw(prefix string) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) BreakEnter(ctx parser.BreakContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) BreakExit(ctx parser.BreakContext) {
	// Get back the last element of the stack
	it := &Break{
		Ctx: ctx,
	}

	l.PushAst(it)
}
