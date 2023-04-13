package ast

import "ogtiger/parser"

type OperationTantQue struct {
	Cond  Ast
	Block Ast
}

func (e OperationTantQue) Display() string {
	return " si"
}

func (l *AstCreatorListener) OperationTantQueEnter(ctx parser.IOperationTantqueContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationTantQueExit(ctx parser.IOperationTantqueContext) {
	OperationTantQue := &OperationTantQue{}

	// Get the first term
	cond := l.AstStack[len(l.AstStack)-1]       // Take it from the top
	l.AstStack = l.AstStack[:len(l.AstStack)-1] // Remove it
	OperationTantQue.Cond = cond                // Store it

	// Get the second term
	blk := l.AstStack[len(l.AstStack)-1]        // Take it from the top
	l.AstStack = l.AstStack[:len(l.AstStack)-1] // Remove it
	OperationTantQue.Block = blk                // Store it

	// Push the new element on the stack
	l.AstStack = append(l.AstStack, OperationTantQue)
}
