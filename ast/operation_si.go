package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz"
)

type OperationSi struct {
	Cond Ast
	Then Ast
	Else Ast
	Ctx  parser.IOperationSiContext
}

func (e *OperationSi) Display() string {
	return " si"
}

func (e *OperationSi) Draw(prefix string, g *graphviz.Graphviz) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) OperationSiEnter(ctx parser.IOperationSiContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationSiExit(ctx parser.IOperationSiContext) {
	OperationSi := &OperationSi{
		Ctx: ctx,
	}

	OperationSi.Cond = l.PopAst()
	OperationSi.Then = l.PopAst()

	if ctx.GetChildCount() == 6 {
		OperationSi.Else = l.PopAst()
	}

	l.PushAst(OperationSi)
}
