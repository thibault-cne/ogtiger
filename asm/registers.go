package asm

type Register string

const (
	R0 Register = "r0"
	R1 Register = "r1"
	R2 Register = "r2"
	R3 Register = "r3"
	R4 Register = "r4"
	R5 Register = "r5"
	R6 Register = "r6"
	R7 Register = "r7"
	R8 Register = "r8"
	R9 Register = "r9"
	R10 Register = "r10"

	BasePointer Register = "r11"
	StackPointer Register = "r13"
	LinkRegister Register = "r14"
	ProgramCounter Register = "r15"
)