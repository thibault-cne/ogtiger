package asm

import (
	"bytes"
	"fmt"
	"os"
)

type Flag string

const (
	EQ Flag = "EQ"
	NE Flag = "NE"
	CS Flag = "CS"
	HS Flag = "HS"
	CC Flag = "CC"
	LO Flag = "LO"
	MI Flag = "MI"
	PL Flag = "PL"
	VS Flag = "VS"
	VC Flag = "VC"
	HI Flag = "HI"
	LS Flag = "LS"
	GE Flag = "GE"
	LT Flag = "LT"
	GT Flag = "GT"
	LE Flag = "LE"
	AL Flag = "AL"
	NI Flag = ""
)

type AssemblyWriter struct {
	indexSID int
	indexQueue []int

	BufferIndex int

	DataBuffer bytes.Buffer
	Buffers []*RegionBuffer
}

type RegionBuffer struct {
	Buffer bytes.Buffer
}

func NewAssemblyWriter() *AssemblyWriter {
	return &AssemblyWriter{
		indexSID: 0,
		indexQueue: []int{},

		BufferIndex: 0,
		DataBuffer: bytes.Buffer{},
		Buffers: []*RegionBuffer{
			{
				Buffer: bytes.Buffer{},
			},
		},
	}
}

func (w *AssemblyWriter) NewRegion()  {
	w.indexSID += 1
	w.indexQueue = append([]int{w.BufferIndex}, w.indexQueue...)
	w.BufferIndex = w.indexSID

	w.Buffers = append(w.Buffers, &RegionBuffer{
		Buffer: bytes.Buffer{},
	})
}

func (w *AssemblyWriter) ExitRegion()  {
	w.BufferIndex, w.indexQueue = w.indexQueue[0], w.indexQueue[1:]
}

func (w *AssemblyWriter) WriteToFile(filename string) {
	f, err := os.Create(fmt.Sprintf("output/%s.s", filename))

	if err != nil {
		panic(err)
	}

	defer f.Close()

	for _, regBuffer := range w.Buffers {
		f.Write(regBuffer.Buffer.Bytes())
		f.WriteString("\n")
	}

	f.WriteString(".data\n")
	f.WriteString("\tformat_str:\t\t\t.ascii\t\t\"%s\\0\"\n")
	f.WriteString("\tformat_int:\t\t\t.ascii\t\t\"%d\\0\"\n")
	f.WriteString("\tformat_debug_int:\t.ascii\t\t\"debug: %d\\n\\0\"\n")
	f.WriteString("\tformat_debug_x:\t\t.ascii\t\t\"debug: %x\\n\\0\"\n")
	f.WriteString("\tformat_debug_addr:\t.ascii\t\t\"debug: %p\\n\\0\"\n")
	f.WriteString("\tscanf_int:\t\t\t.ascii\t\t\"%d\"\n")
	f.Write(w.DataBuffer.Bytes())
}

func (w *AssemblyWriter) Raw(s string) {
	regionBuffer := w.Buffers[w.BufferIndex]
	regionBuffer.Buffer.WriteString(s)
}

func (w *AssemblyWriter) Label(label string) {
	w.Raw(label + ":\n")
}

func (w *AssemblyWriter) SkipLine() {
	w.Raw("\n")
}

func (w *AssemblyWriter) Add(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tADD%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Adc(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tADC%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Sub(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tSUB%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Sbc(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tSBC%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Rsb(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tRSB%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Rsc(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tRSC%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Mov(dst string, val string, flag Flag) {
	instr := fmt.Sprintf("\tMOV%s %s, %s\n", flag, dst, val)

	w.Raw(instr)
}

func (w *AssemblyWriter) Mvn(dst string, val string, flag Flag) {
	instr := fmt.Sprintf("\tMVN%s %s, %s\n", flag, dst, val)

	w.Raw(instr)
}

func (w *AssemblyWriter) And(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tAND%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Eor(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tEOR%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Lsl(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tLSL%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Lsr(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tLSR%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Asr(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tASR%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Ror(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tROR%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Orr(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tORR%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Bic(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tBIC%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Cmp(val1 string, val2 string) {
	instr := fmt.Sprintf("\tCMP %s, %s\n", val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Cmn(val1 string, val2 string) {
	instr := fmt.Sprintf("\tCMN %s, %s\n", val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Tst(val1 string, val2 string) {
	instr := fmt.Sprintf("\tTST %s, %s\n", val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Teq(val1 string, val2 string) {
	instr := fmt.Sprintf("\tTEQ %s, %s\n", val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Ldr(dst string, addr string, flag Flag, offset int) {
	instr := fmt.Sprintf("\tLDR%s %s, [%s, #%d]\n", flag, dst, addr, offset)

	w.Raw(instr)
}

func (w *AssemblyWriter) Ldstr(dst string, addr string, flag Flag) {
	instr := fmt.Sprintf("\tLDR%s %s, =%s\n", flag, dst, addr)

	w.Raw(instr)
}

func (w *AssemblyWriter) Str(dst string, addr string, flag Flag, offset int) {
	instr := fmt.Sprintf("\tSTR%s %s, [%s, #%d]\n", flag, dst, addr, offset)

	w.Raw(instr)
}

func (w *AssemblyWriter) Ldmfd(src string, regs []string) {
	regsStr := ""

	for i, reg := range regs {
		regsStr += reg

		if i != len(regs) - 1 {
			regsStr += ", "
		}
	}


	instr := fmt.Sprintf("\tLDMFD %s!, {%s}\n", src, regsStr)

	w.Raw(instr)
}

func (w *AssemblyWriter) Stmfd(dst string, regs []string) {
	regsStr := ""

	for i, reg := range regs {
		regsStr += reg

		if i != len(regs) - 1 {
			regsStr += ", "
		}
	}


	instr := fmt.Sprintf("\tSTMFD %s!, {%s}\n", dst, regsStr)

	w.Raw(instr)
}

func (w *AssemblyWriter) B(label string, flag Flag) {
	instr := fmt.Sprintf("\tB%s %s\n", flag, label)

	w.Raw(instr)
}

func (w *AssemblyWriter) Bl(label string, flag Flag) {
	instr := fmt.Sprintf("\tBL%s %s\n", flag, label)

	w.Raw(instr)
}

func (w *AssemblyWriter) Fill(addr string, value string) {
	instr := fmt.Sprintf("\t%s FILL %s\n", addr, value)

	w.Raw(instr)
}

func (w *AssemblyWriter) End() {
	w.Raw("\tEND\n")
}

func (w *AssemblyWriter) Comment(comment string, tabs int) {
	instr := ""

	for i := 0; i < tabs; i++ {
		instr += "\t"
	}

	instr = fmt.Sprintf("%s/* %s */\n", instr, comment)

	w.Raw(instr)
}

func (w *AssemblyWriter) Print() {
	instr := `
_print:
	STMFD r13!, {r11, r14}
	MOV r11, r13
	STMFD r13!, {r0, r1}

	LDR r0, [r11, #12]
	LDR r1, [r11, #8]

	BL printf

	MOV r0, #0

	BL fflush

	LDMFD r13!, {r0, r1}
	LDMFD r13!, {r15, r11}
	`

	w.Raw(instr)
}

func (w *AssemblyWriter) Exit(status int) {
	instr := fmt.Sprintf("\tMOV r0, #%d\n", status)
	instr = fmt.Sprintf("_exit:\n%s\tMOV r7, #1\n\tSWI 0\n", instr)

	w.Raw(instr)
}

func (w *AssemblyWriter) AddData(constant *Constants) {
	w.DataBuffer.WriteString(constant.ToASM())
	w.DataBuffer.WriteString("\n")
}
