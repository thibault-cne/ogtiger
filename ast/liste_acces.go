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

type FieldType int

const (
	FieldTypeArr FieldType = iota
	FieldTypeRecord
)

type ListAcces struct {
	Identifiant Ast
	AccesChamps []Ast
	Ctx         parser.ListeAccesContext
	Type        *ttype.TigerType
}

type Field struct {
	Type FieldType
	Access string
	Ctx antlr.ParserRuleContext
}

func (e *ListAcces) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	e.Identifiant.VisitSemControl(slt, L)
	for _, accesChamp := range e.AccesChamps {
		accesChamp.VisitSemControl(slt, L)
	}
	return &e.Ctx
}

func (e *ListAcces) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *ListAcces) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("ListAcces")

	id := e.Identifiant.Draw(g)

	edge, _ := g.CreateEdge("", node, id)
	edge.SetLabel("Identifiant")

	for _, accesChamp := range e.AccesChamps {
		accesChampNode := accesChamp.Draw(g)
		edge, _ = g.CreateEdge("", node, accesChampNode)
		edge.SetLabel("AccesChamp")
	}

	return node
}

func (l *AstCreatorListener) ListAccesEnter(ctx parser.ListeAccesContext) {
	// Nothing to do
}

func (l *AstCreatorListener) ListAccesExit(ctx parser.ListeAccesContext) {
	count, maxCount := 1, ctx.GetChildCount()

	listAcces := &ListAcces{
		Ctx: ctx,
	}

	acces := make([]Field, 0)
	// Get the accesChamps
	for count < maxCount {
		text := ctx.GetChild(count).(*antlr.TerminalNodeImpl).GetText()
		fieldAccessed := l.PopAst()

		if text == "." {
			count += 1
			val, c := GetChildTextAndCtx(ctx.GetChild(count))

			f := Field{
				Type: FieldTypeRecord,
				Access: val,
				Ctx: c,
			}
			acces = append([]Field{ f }, acces...)
		} else {
			count += 2
			val, c := GetChildTextAndCtx(ctx.GetChild(count - 1))

			f := Field{
				Type: FieldTypeArr,
				Access: val,
				Ctx: c,
			}
			acces = append([]Field{ f }, acces...)

			if !fieldAccessed.ReturnType().Equals(ttype.NewTigerType(ttype.Int)) {
				l.Logger.NewSemanticError(logger.ErrorArrayIndexIsNotAnInteger, f.Ctx, f.Access)
			}

		}

		// Prepare the accesChamp
		listAcces.AccesChamps = append([]Ast{ fieldAccessed }, listAcces.AccesChamps...)

		count += 1
	}

	// Get the identifiant
	listAcces.Identifiant = l.PopAst()

	// Set the type
	t := listAcces.Identifiant.ReturnType()
	element := listAcces.Identifiant.(*Identifiant).Id

	for _, a := range acces {
		if a.Type == FieldTypeArr {
			if t.ElementType == nil {
				l.Logger.NewSemanticError(logger.ErrorVarIsNotAnArray, a.Ctx, element)
				break
			}

			element = fmt.Sprintf("%s%s", element, a.Access)
			t = t.ElementType

		} else {
			tmp, err := t.GetRecordFieldType(a.Access)

			if err != nil {
				l.Logger.NewSemanticError(logger.ErrorFieldDoesNotExist, a.Ctx, a.Access, t)
				break
			}

			element = fmt.Sprintf("%s.%s", element, a.Access)
			t = tmp
		}
	}

	listAcces.Type = t

	// Push the listAcces
	l.PushAst(listAcces)
}
