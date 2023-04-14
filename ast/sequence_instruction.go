package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type SequenceInstruction struct {
	Instructions []Ast
	Ctx          parser.ISequenceInstructionContext
	Type         ttype.TigerType
}

func (e *SequenceInstruction) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *SequenceInstruction) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("Sequence")

	for _, instruction := range e.Instructions {
		instructionNode := instruction.Draw(g)
		g.CreateEdge("Instruction", node, instructionNode)
	}

	return node
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

	// Reverse the slice
	for i := len(opSeq.Instructions)/2 - 1; i >= 0; i-- {
		opp := len(opSeq.Instructions) - 1 - i
		opSeq.Instructions[i], opSeq.Instructions[opp] = opSeq.Instructions[opp], opSeq.Instructions[i]
	}

	l.PushAst(opSeq)
}
