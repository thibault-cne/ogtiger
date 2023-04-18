package ast

import (
	"fmt"
	"ogtiger/logger"
	"ogtiger/parser"
	"ogtiger/slt"
	"ogtiger/ttype"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-graphviz/cgraph"
)

type Identifiant struct {
	Id   string
	Ctx  parser.IIdentifiantContext
	Type *ttype.TigerType
}

func (e *Identifiant) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	return e.Ctx
}

func (e *Identifiant) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *Identifiant) GetErrorCount() int {
	return 0
}

func (e *Identifiant) Draw(g *cgraph.Graph) *cgraph.Node {
	id := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(id)
	node.SetLabel(e.Id)

	return node
}

func (l *AstCreatorListener) IdentifiantEnter(ctx parser.IIdentifiantContext) {
	// Nothing to do
}

func (l *AstCreatorListener) IdentifiantExit(ctx parser.IIdentifiantContext) {
	identifiant := &Identifiant{
		Id:  ctx.GetText(),
		Ctx: ctx,
	}

	t, _ := l.Slt.GetSymbolType(ctx.GetText())

	identifiant.Type = t

	l.PushAst(identifiant)
}
