package ast

import "ogtiger/parser"

type OperationSi struct {
	Cond Ast
	Then Ast
	Else Ast
}

func (e OperationSi) Display() string {
	return " si"
}

func (l *AstCreatorListener) OperationSiEnter(ctx parser.IOperationSiContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationSiExit(ctx parser.IOperationSiContext) {
	OperationSi := &OperationSi{}

	// Get the first term
	cond := l.AstStack[len(l.AstStack)-1]       // Take it from the top
	l.AstStack = l.AstStack[:len(l.AstStack)-1] // Remove it
	OperationSi.Cond = cond                     // Store it

	// Get the second term
	then := l.AstStack[len(l.AstStack)-1]       // Take it from the top
	l.AstStack = l.AstStack[:len(l.AstStack)-1] // Remove it
	OperationSi.Cond = then                     // Store it

	if ctx.GetChildCount() == 6 {
		// Get the third term
		el := l.AstStack[len(l.AstStack)-1]         // Take it from the top
		l.AstStack = l.AstStack[:len(l.AstStack)-1] // Remove it
		OperationSi.Cond = el                       // Store it
	}

	// Push the new element on the stack
	l.AstStack = append(l.AstStack, OperationSi)
}
