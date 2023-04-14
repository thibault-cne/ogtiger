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
	return &SymbolTable{
		Table:  make(map[string]*Symbol),
		Region: 0,
		Scope:  0,
	}
}

func (s *SymbolTable) AddSymbol(name string, symbol *Symbol) {
	s.Table[name] = symbol
}

func (s *SymbolTable) GetSymbol(name string) *Symbol {
	if val, ok := s.Table[name]; ok {
		return val
	}

	if s.Parent != nil {
		return s.Parent.GetSymbol(name)
	}

	return nil
}

func (s *SymbolTable) GetSymbolType(name string) (*ttype.TigerType, error) {
	if val, ok := s.Table[name]; ok {
		return &val.Type, nil
	}

	if s.Parent != nil {
		return s.Parent.GetSymbolType(name)
	}

	return nil, fmt.Errorf("Symbol %s not found", name)
}
