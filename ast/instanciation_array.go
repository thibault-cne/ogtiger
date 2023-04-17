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

type InstanciationArray struct {
	Id   Ast
	Size Ast
	DefaultValue Ast
	Ctx  parser.InstanciationArrayContext
	Type *ttype.TigerType
}

func (e *InstanciationArray) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	e.Id.VisitSemControl(slt, L)
	e.Size.VisitSemControl(slt, L)
	e.DefaultValue.VisitSemControl(slt, L)


	return &e.Ctx
}

func (e *InstanciationArray) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *InstanciationArray) Draw(g *cgraph.Graph) *cgraph.Node {
	id := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(id)
	node.SetLabel("InstanciationArray")

	idNode := e.Id.Draw(g)
	g.CreateEdge("Id", node, idNode)

	sizeNode := e.Size.Draw(g)
	g.CreateEdge("Size", node, sizeNode)

	defaultValueNode := e.DefaultValue.Draw(g)
	g.CreateEdge("DefaultValue", node, defaultValueNode)

	return node
}

func (l *AstCreatorListener) InstanciationArrayEnter(ctx parser.InstanciationArrayContext) {
	// Nothing to do
}

func (l *AstCreatorListener) InstanciationArrayExit(ctx parser.InstanciationArrayContext) {
	instanciationArray := &InstanciationArray{
		Ctx:  ctx,
	}

	instanciationArray.DefaultValue = l.PopAst()
	instanciationArray.Size = l.PopAst()
	instanciationArray.Id = l.PopAst()

	instanciationArray.Type = ttype.NewArrayType(instanciationArray.DefaultValue.ReturnType())

	l.PushAst(instanciationArray)
}
