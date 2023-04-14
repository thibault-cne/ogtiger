package ast

import (
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type DeclarationTypeClassique struct {
	Identifiant Ast
	TType       Ast
	Ctx         parser.DeclarationTypeClassiqueContext
	Type        ttype.TigerType
}

func (e *DeclarationTypeClassique) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *DeclarationTypeClassique) Display() string {
	return " alias"
}

func (e *DeclarationTypeClassique) Draw(prefix string, g *cgraph.Graph) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) DeclarationTypeClassiqueEnter(ctx parser.DeclarationTypeClassiqueContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) DeclarationTypeClassiqueExit(ctx parser.DeclarationTypeClassiqueContext) {
	// Get back the last element of the stack
	declType := &DeclarationTypeClassique{
		Ctx: ctx,
	}

	declType.Identifiant = l.PopAst()
	declType.TType = l.PopAst()

	l.PushAst(declType)
}
