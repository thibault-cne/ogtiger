package ast

import (
	"ogtiger/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type OperationMultiplication struct {
	Left  Ast
	Right []*OperationMultiplicationFD
}

type OperationMultiplicationFD struct {
	Op    string
	Right Ast
}

func (e *OperationMultiplication) Display() string {
	return " addition"
}

func (l *AstCreatorListener) OperationMultiplicationEnter(ctx parser.ExpressionContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) OperationMultiplicationExit(ctx parser.ExpressionContext) {
	// Get back the last element of the stack
	opAddition := &OperationMultiplication{}

	if ctx.GetChildCount() == 1 {
		return
	}

	// Get the first term
	left := l.AstStack[len(l.AstStack)-1]       // Take it from the top
	l.AstStack = l.AstStack[:len(l.AstStack)-1] // Remove it
	opAddition.Left = left                      // Store it

	// Get minus and plus and term number
	for i := 0; i < (ctx.GetChildCount()-1)/2; i++ {
		right := &OperationMultiplicationFD{}

		right.Op = ctx.GetChild(2*i + 1).(*antlr.TerminalNodeImpl).GetText()
		right.Right = l.AstStack[len(l.AstStack)-1]

		// Pop the last element of the stack
		l.AstStack = l.AstStack[:len(l.AstStack)-1]

		opAddition.Right = append(opAddition.Right, right)
	}

	// Push the new element on the stack
	l.AstStack = append(l.AstStack, opAddition)
}
