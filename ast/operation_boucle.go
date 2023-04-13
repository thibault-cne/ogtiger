package ast

import "ogtiger/parser"

type OperationBoucle struct {
	Start    Ast
	StartVal Ast
	EndVal   Ast
	Block    Ast
	Ctx      parser.IOperationBoucleContext
}

func (e *OperationBoucle) Display() string {
	return " pour"
}

func (e *OperationBoucle) Draw(prefix string) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) OperationBoucleEnter(ctx parser.IOperationBoucleContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationBoucleExit(ctx parser.IOperationBoucleContext) {
	oB := &OperationBoucle{
		Ctx: ctx,
	}

	oB.Start = l.PopAst()
	oB.StartVal = l.PopAst()
	oB.EndVal = l.PopAst()
	oB.Block = l.PopAst()

	// Push the new element on the stack
	l.AstStack = append(l.AstStack, oB)
}
