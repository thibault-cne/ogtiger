package ast

import (
	"ogtiger/parser"
)

type SequenceInstruction struct {
	Instructions []Ast
}

func (e *SequenceInstruction) Display() string {
	return " addition"
}

func (l *AstCreatorListener) SequenceInstructionEnter(ctx parser.ExpressionContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) SequenceInstructionExit(ctx parser.ExpressionContext) {
	// Get back the last element of the stack
	opSeq := &SequenceInstruction{}

	// Get minus and plus and term number
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		child := l.AstStack[len(l.AstStack)-1]
		// Pop the last element of the stack
		l.AstStack = l.AstStack[:len(l.AstStack)-1]

		opSeq.Instructions = append(opSeq.Instructions, child)
	}

	// We have only one instruction, we don't need a sequence
	if len(opSeq.Instructions) == 1 {
		l.AstStack = append(l.AstStack, opSeq.Instructions[0])
		return
	}

	// Push the new element on the stack
	l.AstStack = append(l.AstStack, opSeq)
}
