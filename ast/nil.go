package ast

import (
	"ogtiger/parser"
)

type Nil struct {
	Ctx parser.NilContext
}

func (e *Nil) Display() string {
	return " nil"
}

func (l *AstCreatorListener) NilEnter(ctx parser.NilContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) NilExit(ctx parser.NilContext) {
	// Get back the last element of the stack
	it := &Nil{
		Ctx: ctx,
	}

	l.PushAst(it)
}
