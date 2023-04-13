package ast

import (
	"ogtiger/parser"
)

type Definition struct {
	Declarations []Ast
	Expressions  []Ast
}

func (e *Definition) Display() string {
	return " definition"
}

func (l *AstCreatorListener) DefinitionEnter(ctx parser.DefinitionContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) DefinitionExit(ctx parser.DefinitionContext) {
	// Get back the last element of the stack
	expr := &Definition{}

	for range ctx.AllDeclaration() {
		expr.Declarations = append(expr.Declarations, l.PopAst())
	}

	for range ctx.AllExpression() {
		expr.Expressions = append(expr.Expressions, l.PopAst())
	}

	l.PushAst(expr)
}
