package ast

import (
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type OperationNegation struct {
	Expr Ast
	Ctx  parser.IOperationNegationContext
	Type ttype.TigerType
}

func (e *OperationNegation) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *OperationNegation) Display() string {
	return " negation"
}

func (e *OperationNegation) Draw(prefix string, g *cgraph.Graph) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) OperationNegationEnter(ctx parser.IOperationNegationContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationNegationExit(ctx parser.IOperationNegationContext) {
	operationNegation := &OperationNegation{
		Ctx: ctx,
	}

	operationNegation.Expr = l.PopAst()

	l.PushAst(operationNegation)
}
