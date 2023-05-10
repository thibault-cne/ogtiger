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

	// Check if the size is an integer
	if !instanciationArray.Size.ReturnType().Equals(ttype.NewTigerType(ttype.Int)) {
		value, c := GetChildTextAndCtx(ctx.GetChild(2))

		l.Logger.NewSemanticError(logger.ErrorArraySizeIsNotAnInteger, c, value)
	}

	if !instanciationArray.Id.ReturnType().ElementType.Equals(instanciationArray.DefaultValue.ReturnType()) {
		typeName := instanciationArray.Id.(*Identifiant).Id
		actualType := instanciationArray.DefaultValue.ReturnType()
		expectedType := instanciationArray.Id.ReturnType().ElementType

		_, c := GetChildTextAndCtx(ctx.GetChild(5))
		
		l.Logger.NewSemanticError(logger.ErrorWrongDefaultValueArray, c, typeName, expectedType, actualType)
	}

	instanciationArray.Type = ttype.NewArrayType(instanciationArray.DefaultValue.ReturnType())

	l.PushAst(instanciationArray)
}

func (e *InstanciationArray) EnterAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	defer e.ExitAsm(writer, slt)
}

func (e *InstanciationArray) ExitAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	// Nothing to do
}
