package ast

import (
	"ogtiger/parser"
)

type SequenceInstruction struct {
	Instructions []Ast
	Ctx          parser.ISequenceInstructionContext
}

func (e *SequenceInstruction) Display() string {
	return " sequence"
}

func (l *AstCreatorListener) SequenceInstructionEnter(ctx parser.ISequenceInstructionContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) SequenceInstructionExit(ctx parser.ISequenceInstructionContext) {
	// Get back the last element of the stack
	opSeq := &SequenceInstruction{
		Ctx: ctx,
	}

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
