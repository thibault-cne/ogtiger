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
	Type         *ttype.TigerType
}

func (e *SequenceInstruction) ReturnType() *ttype.TigerType {
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

func (l *SemControlListener) SequenceInstructionEnter(ctx parser.ISequenceInstructionContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *SemControlListener) SequenceInstructionExit(ctx parser.ISequenceInstructionContext) {
	// Get back the last element of the stack
	opSeq := &SequenceInstruction{
		Ctx: ctx,
	}

	// Get minus and plus and term number
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		opSeq.Instructions = append([]Ast{l.PopAst()}, opSeq.Instructions...)
	}

	// Set the type of the expression
	opSeq.Type = opSeq.Instructions[len(opSeq.Instructions)-1].ReturnType()

	l.PushAst(opSeq)
}
