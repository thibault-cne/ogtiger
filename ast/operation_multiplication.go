package ast

import (
	"fmt"
	"ogtiger/asm"
	"ogtiger/logger"
	"ogtiger/parser"
	"ogtiger/slt"
	"ogtiger/ttype"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-graphviz/cgraph"
)

type OperationMultiplication struct {
	Left  Ast
	Right Ast
	Ctx   parser.IOperationMultiplicationContext
	Type  *ttype.TigerType
}

func (e *OperationMultiplication) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	e.Left.VisitSemControl(slt, L)
	e.Right.VisitSemControl(slt, L)

	typeInt := ttype.NewTigerType(ttype.Int)
	typeString := ttype.NewTigerType(ttype.String)

	if e.Left.ReturnType().Equals(e.Right.ReturnType()) && !e.Left.ReturnType().Equals(typeInt) {
		if !((e.Left.ReturnType().Equals(typeString) && e.Right.ReturnType().Equals(typeInt)) || (e.Left.ReturnType().Equals(typeInt) && e.Right.ReturnType().Equals(typeString))) {
			L.NewSemanticError("Multiplication between non integer and non string", e.Ctx)
		}
	}

	return e.Ctx
}

type OperationDivision struct {
	Left  Ast
	Right Ast
	Ctx   parser.IOperationMultiplicationContext
	Type  *ttype.TigerType
}

func (e *OperationDivision) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	leftCtx := e.Left.VisitSemControl(slt, L)
	rightCtx := e.Right.VisitSemControl(slt, L)

	if !e.Left.ReturnType().Equals(ttype.NewTigerType(ttype.Int)) {
		L.NewSemanticError("Left side of the division is not an integer", leftCtx)
	}

	if !e.Right.ReturnType().Equals(ttype.NewTigerType(ttype.Int)) {
		L.NewSemanticError("Right side of the division is not an integer", rightCtx)
	}

	return e.Ctx
}

func (e *OperationMultiplication) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *OperationDivision) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *OperationMultiplication) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("*")

	left := e.Left.Draw(g)
	g.CreateEdge("Left", node, left)
	right := e.Right.Draw(g)
	g.CreateEdge("Right", node, right)

	return node
}

func (e *OperationDivision) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("/")

	left := e.Left.Draw(g)
	g.CreateEdge("Left", node, left)
	right := e.Right.Draw(g)
	g.CreateEdge("Right", node, right)

	return node
}

func (l *AstCreatorListener) OperationMultiplicationEnter(ctx parser.IOperationMultiplicationContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) OperationMultiplicationExit(ctx parser.IOperationMultiplicationContext) {
	if ctx.GetChildCount() == 1 {
		return
	}

	// Get back the elements needed from the stack
	elements := make([]Ast, 0)

	for i := 0; i < (ctx.GetChildCount()+1)/2; i++ {
		elements = append(elements, l.PopAst())
	}

	node := elements[len(elements)-1]
	elements = elements[:len(elements)-1]

	// TODO: Check if the type is correct

	// Get minus and plus and term number
	for i := 0; 2*i < (ctx.GetChildCount() - 1); i++ {
		switch ctx.GetChild(2*i + 1).(*antlr.TerminalNodeImpl).GetText() {
		case "*":
			temp := &OperationMultiplication{
				Ctx:   ctx,
				Left:  node,
				Right: elements[len(elements)-1],
				Type:  ttype.NewTigerType(ttype.Int),
			}
			node = temp
		case "/":
			temp := &OperationDivision{
				Ctx:   ctx,
				Left:  node,
				Right: elements[len(elements)-1],
				Type:  ttype.NewTigerType(ttype.Int),
			}
			node = temp
		}

		elements = elements[:len(elements)-1]
	}

	l.PushAst(node)
}

func (e *OperationMultiplication) EnterAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	defer e.ExitAsm(writer, slt)

	writer.SkipLine()
	writer.Comment("Multiplication", 1)
	
	e.Left.EnterAsm(writer, slt)

	writer.Stmfd(string(asm.StackPointer), []string{string(asm.R8)})

	e.Right.EnterAsm(writer, slt)

	writer.Ldmfd(string(asm.StackPointer), []string{string(asm.R0)})
	writer.Mul(string(asm.R8), string(asm.R0), string(asm.R8), asm.NI)
}

func (e *OperationMultiplication) ExitAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	// Nothing to do
}

func (e *OperationDivision) EnterAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	defer e.ExitAsm(writer, slt)

	writer.SkipLine()
	writer.Comment("Divide", 1)
	
	e.Left.EnterAsm(writer, slt)

	writer.Stmfd(string(asm.StackPointer), []string{string(asm.R8)})

	e.Right.EnterAsm(writer, slt)

	writer.Ldmfd(string(asm.StackPointer), []string{string(asm.R0)})
	writer.Sdiv(string(asm.R8), string(asm.R0), string(asm.R8), asm.NI)
}

func (e *OperationDivision) ExitAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	// Nothing to do
}
