package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz/cgraph"
)

type Program struct {
	Expr Ast
	Ctx  parser.IProgramContext
}

func (p *Program) Display() string {
	return " program"
}

func (e *Program) Draw(prefix string, g *cgraph.Graph) {
	// TODO: Draw the AST
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
