package ast

import (
	"ogtiger/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Integer struct {
	Valeur string
	Ctx    parser.EntierContext
}

func (e *Integer) Display() string {
	return " alias"
}

func (l *AstCreatorListener) IntegerEnter(ctx parser.EntierContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) IntegerExit(ctx parser.EntierContext) {
	// Get back the last element of the stack
	it := &Integer{
		Ctx: ctx,
	}

	it.Valeur = ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetText()

	l.PushAst(it)
}
