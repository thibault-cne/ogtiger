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
	return " multiplication"
}

func (l *AstCreatorListener) OperationMultiplicationEnter(ctx parser.ExpressionContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) OperationMultiplicationExit(ctx parser.ExpressionContext) {
	// Get back the last element of the stack
	opMultiplication := &OperationMultiplication{}

	if ctx.GetChildCount() == 1 {
		return
	}

	opMultiplication.Left = l.PopAst()

	// Get minus and plus and term number
	for i := 0; i < (ctx.GetChildCount()-1)/2; i++ {
		right := &OperationMultiplicationFD{}

		right.Op = ctx.GetChild(2*i + 1).(*antlr.TerminalNodeImpl).GetText()
		right.Right = l.PopAst()

		opMultiplication.Right = append(opMultiplication.Right, right)
	}

	l.PushAst(opMultiplication)
}
