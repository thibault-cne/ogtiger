package asm

import "fmt"

var SID = 0

type Constants struct {
	Id string

	StrValue string
	IntValue int

	Type int
}

func NewIntConstant(intValue int) *Constants {
	SID += 1

	return &Constants{
		Id: fmt.Sprintf("%d", SID),
		StrValue: "",
		IntValue: intValue,
		Type: 0,
	}
}

func NewStrConstant(strValue string) *Constants {
	SID += 1

	return &Constants{
		Id: fmt.Sprintf("%d", SID),
		StrValue: strValue,
		IntValue: 0,
		Type: 1,
	}
}

func (c *Constants) ToASM() string {
	switch c.Type {
	case 0:
		return fmt.Sprintf("\tstring_%s:\t\t\t.word\t\t%d", c.Id, c.IntValue)
	case 1:
		return fmt.Sprintf("\tstring_%s:\t\t\t.ascii\t\t\"%s\\0\"", c.Id, c.StrValue)
	}

	return ""
}

func (c *Constants) GetId() string {
	return fmt.Sprintf("string_%s", c.Id)
}