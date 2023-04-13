package ast

import (
	"ogtiger/parser"
)

type Program struct {
	Expr Ast
	Ctx  parser.IProgramContext
}

func (p *Program) Display() string {
	return " program"
}

func (l *AstCreatorListener) ProgramEnter(ctx parser.IProgramContext) {
	// l.AstStack = append(l.AstStack, &Program{})
}

func (l *AstCreatorListener) ProgramExit(ctx parser.IProgramContext) {
	prog := &Program{
		Ctx: ctx,
	}

	prog.Expr = l.PopAst()

	// Push the new element on the stack
	l.PushAst(prog)
}
