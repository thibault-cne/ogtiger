package slt

import (
	"fmt"
	"ogtiger/ttype"
)

type SymbolTable struct {
	Table  map[string]*Symbol
	Region int
	Scope  int
	Parent *SymbolTable
}

func NewSymbolTable() *SymbolTable {
	slt := &SymbolTable{
		Table:  make(map[string]*Symbol),
		Region: 0,
		Scope:  0,
	}

	// Add the built-in functions and types
	slt.CreateSymbol("print", ttype.NewFunctionType(ttype.NewTigerType(ttype.NoReturn), []*ttype.FunctionParameter{}))
	slt.CreateSymbol("flush", ttype.NewFunctionType(ttype.NewTigerType(ttype.NoReturn), []*ttype.FunctionParameter{}))
	slt.CreateSymbol("getchar", ttype.NewFunctionType(ttype.NewTigerType(ttype.String), []*ttype.FunctionParameter{}))

	slt.CreateSymbol("int", ttype.NewTigerType(ttype.Int))
	slt.CreateSymbol("string", ttype.NewTigerType(ttype.String))
	return slt
}

func (s *SymbolTable) CreateSymbol(name string, t *ttype.TigerType) {
	s.Table[name] = &Symbol{
		Name: name,
		Type: t,
	}
}

func (s *SymbolTable) AddSymbol(name string, symbol *Symbol) {
	s.Table[name] = symbol
}

func (s *SymbolTable) GetSymbol(name string) (*Symbol, error) {
	if val, ok := s.Table[name]; ok {
		return val, nil
	}

	if s.Parent != nil {
		return s.Parent.GetSymbol(name)
	}

	return nil, fmt.Errorf("Symbol %s not found", name)
}

func (s *SymbolTable) GetSymbolInScoope(name string) (*Symbol, error) {
	if val, ok := s.Table[name]; ok {
		return val, nil
	}

	return nil, fmt.Errorf("Symbol %s not found", name)
}

func (s *SymbolTable) GetSymbolType(name string) (*ttype.TigerType, error) {
	if val, ok := s.Table[name]; ok {
		return val.Type, nil
	}

	if s.Parent != nil {
		return s.Parent.GetSymbolType(name)
	}

	return nil, fmt.Errorf("Symbol %s not found", name)
}
