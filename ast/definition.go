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

type Definition struct {
	Declarations []Ast
	Expressions  []Ast
	Slt          *slt.SymbolTable
	Ctx          parser.IDefinitionContext
	Type         *ttype.TigerType
	DeclarationCount int
}

func (e *Definition) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	for _, declaration := range e.Declarations {
		declaration.VisitSemControl(e.Slt, L)
	}

	for _, expression := range e.Expressions {
		expression.VisitSemControl(e.Slt, L)
	}

	return e.Ctx
}

func (e *Definition) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *Definition) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("Definition")

	for _, declaration := range e.Declarations {
		declarationNode := declaration.Draw(g)
		g.CreateEdge("", node, declarationNode)
	}

	for _, expression := range e.Expressions {
		expressionNode := expression.Draw(g)
		g.CreateEdge("", node, expressionNode)
	}

	return node
}

func (l *AstCreatorListener) DefinitionEnter(ctx parser.IDefinitionContext) {
	// Create a TDS for the definition
	l.Slt = l.Slt.CreateRegion()
}

func (l *AstCreatorListener) DefinitionExit(ctx parser.IDefinitionContext) {
	expr := &Definition{
		Ctx: ctx,
		DeclarationCount: 0,
	}

	for range ctx.AllExpression() {
		// Prepend the expressions to the list
		expr.Expressions = append([]Ast{ l.PopAst() }, expr.Expressions...)
	}

	expr.Type = expr.Expressions[len(expr.Expressions)-1].ReturnType()

	for range ctx.AllDeclaration() {
		// Prepend the declarations to the list
		expr.Declarations = append([]Ast{ l.PopAst() }, expr.Declarations...)

		switch expr.Declarations[0].(type) {
		case *DeclarationFontion:
			break
		default:
			expr.DeclarationCount += 1
		}
	}

	// Pop the TDS
	expr.Slt = l.Slt
	l.Slt = l.Slt.Parent

	l.PushAst(expr)
}

func (e *Definition) EnterAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	defer e.ExitAsm(writer, slt)

	writer.NewRegion()

	label := fmt.Sprintf("blk_%d_%d", e.Slt.Region, e.Slt.Scope)
	writer.Label(label)
	registers := []string{string(asm.BasePointer), string(asm.LinkRegister), string(asm.R0)}

	writer.Comment("Display edit on entering a block", 1)
	writer.Ldr("r0", "r10", asm.NI, - (e.Slt.Scope * 4))
	writer.Stmfd(string(asm.StackPointer), registers)
	writer.Mov(string(asm.BasePointer), string(asm.StackPointer), asm.NI)
	writer.Str(string(asm.BasePointer), "r10", asm.NI, - (e.Slt.Scope * 4))

	for _, a := range e.Declarations {
		a.EnterAsm(writer, e.Slt)
	}

	for _, a := range e.Expressions {
		a.EnterAsm(writer, e.Slt)
	}
}

func (e *Definition) ExitAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	writer.SkipLine()
	writer.Comment("Remove declaration from stack", 1)
	writer.Add(string(asm.StackPointer), string(asm.StackPointer), fmt.Sprintf("#%d", e.DeclarationCount * 4), asm.NI)
	
	writer.SkipLine()
	writer.Comment("Display edit on exiting a block", 1)
	writer.Ldmfd(string(asm.StackPointer), []string{string(asm.R0)})
	writer.Str("r0", "r10", asm.NI, - (e.Slt.Scope * 4))

	registers := []string{string(asm.ProgramCounter), string(asm.BasePointer)}
	writer.Ldmfd(string(asm.StackPointer), registers)

	writer.ExitRegion()
}
