package slt

import "ogtiger/ttype"

type Symbol struct {
	Name   string
	Type   *ttype.TigerType
	Offset int
}
