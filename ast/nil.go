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

type Nil struct {
	Ctx  parser.NilContext
	Type *ttype.TigerType
}

func (e *Nil) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	return &e.Ctx
}

func (e *Nil) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *Nil) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("Nil")

	return node
}

func (l *AstCreatorListener) NilEnter(ctx parser.NilContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) NilExit(ctx parser.NilContext) {
	// Get back the last element of the stack
	it := &Nil{
		Ctx:  ctx,
		Type: ttype.NewTigerType(ttype.AnyRecord),
	}

	l.PushAst(it)
}

func (e *Nil) EnterAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	defer e.ExitAsm(writer, slt)
}

func (e *Nil) ExitAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	// Nothing to do
}
