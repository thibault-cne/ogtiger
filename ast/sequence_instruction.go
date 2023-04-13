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
		opSeq.Instructions = append(opSeq.Instructions, l.PopAst())
	}

	// We have only one instruction, we don't need a sequence
	if len(opSeq.Instructions) == 1 {
		l.PushAst(opSeq.Instructions[0])
		return
	}

	l.PushAst(opSeq)
}
