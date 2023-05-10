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

func (e *AppelFonction) EnterAsm(writer *asm.AssemblyWriter) {
	defer e.ExitAsm(writer)

	if e.Identifiant.(*Identifiant).Id == "print" {
		writer.Ldr("R0", "format_str", asm.NI, 0)
		writer.Stmfd("SP", []string{"R0"})
	}

	for _, a := range e.Args {
		a.EnterAsm(writer)
	}
}

func (e *AppelFonction) ExitAsm(writer *asm.AssemblyWriter) {
	// Nothing to do
}
