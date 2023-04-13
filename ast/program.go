package ast

import (
	"ogtiger/parser"
)

type Program struct {
	Expr Ast
}

func (p *Program) Display() string {
	return " program"
}

func (l *AstCreatorListener) ProgramEnter(ctx parser.IProgramContext) {
	// l.AstStack = append(l.AstStack, &Program{})
}

func (l *AstCreatorListener) ProgramExit(ctx parser.IProgramContext) {
	// We push the Program struct to the stack as it is the entrypoint

	prog := &Program{}

	prog.Expr = l.PopAst()

	// Push the new element on the stack
	l.PushAst(prog)
}
