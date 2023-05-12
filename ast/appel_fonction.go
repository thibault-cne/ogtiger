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

type AppelFonction struct {
	Identifiant Ast
	Args        []Ast
	Ctx         parser.AppelFonctionContext
	Type        *ttype.TigerType
}

func (e *AppelFonction) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	idCtx := e.Identifiant.VisitSemControl(slt, L)
	id := e.Identifiant.(*Identifiant).Id
	f, err := slt.GetSymbol(id)
	if err != nil || f.Type.ID != ttype.Function {
		L.NewSemanticError(logger.ErrorFunctionNotDefined, idCtx, id)
		return &e.Ctx
	}

	if len(e.Args) != len(f.Type.Parameters) {
		L.NewSemanticError(logger.ErrorWrongNumberOfArgs, &e.Ctx, id, len(f.Type.Parameters), len(e.Args))
	} else {
		for i, arg := range e.Args {
			argCtx := arg.VisitSemControl(slt, L)

			if !arg.ReturnType().Equals(f.Type.Parameters[i].Type) {
				L.NewSemanticError(logger.ErrorWrongTypeOfArgs, argCtx, id, f.Type.Parameters[i].Type, f.Type.Parameters[i].Name, arg.ReturnType())
			}
		}
	}

	return &e.Ctx
}

func (e *AppelFonction) ReturnType() *ttype.TigerType {
	return e.Type
}

func (a *AppelFonction) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", a)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("AppelFonction")

	id := a.Identifiant.Draw(g)
	g.CreateEdge("Id", node, id)

	nodeAId := fmt.Sprintf("NA%p", a.Args)
	nodeA, _ := g.CreateNode(nodeAId)
	nodeA.SetLabel("ArgFonction")

	g.CreateEdge("Args", node, nodeA)

	for _, arg := range a.Args {
		argNode := arg.Draw(g)
		g.CreateEdge("", nodeA, argNode)
	}

	return node
}

func (l *AstCreatorListener) AppelFonctionEnter(ctx parser.AppelFonctionContext) {
	// Nothing to do
}

func (l *AstCreatorListener) AppelFonctionExit(ctx parser.AppelFonctionContext) {
	appelFonction := &AppelFonction{
		Ctx: ctx,
	}

	// Get the args
	for i := 0; i < len(ctx.AllExpression()); i++ {
		// Pop the last element of the stack
		// Prepend it to the args
		appelFonction.Args = append([]Ast{l.PopAst()}, appelFonction.Args...)
	}

	// Get the identifiant
	appelFonction.Identifiant = l.PopAst()

	// Set the type
	appelFonction.Type = appelFonction.Identifiant.ReturnType()

	// Push the new element on the stack
	l.PushAst(appelFonction)
}

func (e *AppelFonction) EnterAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	defer e.ExitAsm(writer, slt)

	fnId := e.Identifiant.(*Identifiant).Id

	writer.Comment(fmt.Sprintf("Call the %s fn", fnId), 1)

	if fnId == "print" {
		writer.Comment("Load the format string parameter for the print function", 1)
		writer.Ldstr("r0", "format_debug_int", asm.NI)
		writer.Stmfd(string(asm.StackPointer), []string{"r0"})
		writer.SkipLine()
	}

	for i, a := range e.Args {
		a.EnterAsm(writer, slt)

		// Result is stored in R8 so we just push it in the stack
		if i != 0 {
			writer.SkipLine()
		}
		writer.Comment(fmt.Sprintf("Arg %s of function %s", e.Type.Parameters[i].Name, fnId), 1)
		writer.Stmfd(string(asm.StackPointer), []string{"r8"})
	}

	writer.SkipLine()
	writer.Bl(fmt.Sprintf("_%s", fnId), asm.NI)
}

func (e *AppelFonction) ExitAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	writer.SkipLine()

	writer.Comment("Remove args", 1)
	if e.Identifiant.(*Identifiant).Id == "print" {
		writer.Add(string(asm.StackPointer), string(asm.StackPointer), "#8", asm.NI)
	} else {
		writer.Add(string(asm.StackPointer), string(asm.StackPointer), fmt.Sprintf("#%d", len(e.Type.Parameters) * 4), asm.NI)
	}
}
