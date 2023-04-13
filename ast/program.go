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

func  (l *AstCreatorListener) ExprEnter(ctx parser.IProgramContext) {
	// l.AstStack = append(l.AstStack, &Program{})
}

func  (l *AstCreatorListener) ExprExit(ctx parser.IProgramContext) {
	// We push the Program struct to the stack as it is the entrypoint

	prog := &Program{}

	// Unstack the last expr and put it in the struct
	expr := l.AstStack[len(l.AstStack)-1]

	// Pop the last element of the stack
	l.AstStack = l.AstStack[:len(l.AstStack)-1]

	prog.Expr = expr

	// Push the new element on the stack
	l.AstStack = append(l.AstStack, prog)
}