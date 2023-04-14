package ast

import (
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type OperationSi struct {
	Cond Ast
	Then Ast
	Else Ast
	Ctx  parser.IOperationSiContext
	Type ttype.TigerType
}

func (e *OperationSi) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *OperationSi) Display() string {
	return " si"
}

func (e *OperationSi) Draw(prefix string, g *cgraph.Graph) {
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
