// Code generated from ./grammar/tiger.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // tiger

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type tigerParser struct {
	*antlr.BaseParser
}

var tigerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tigerParserInit() {
	staticData := &tigerParserStaticData
	staticData.literalNames = []string{
		"", "':='", "'|'", "'&'", "'='", "'<>'", "'<'", "'>='", "'>'", "'<='",
		"'+'", "'-'", "'*'", "'/'", "'('", "';'", "')'", "','", "'['", "']'",
		"'of'", "'.'", "'{'", "'}'", "'if'", "'then'", "'else'", "'while'",
		"'do'", "'for'", "'to'", "'let'", "'in'", "'end'", "'type'", "'array'",
		"':'", "'function'", "'var'", "'nil'", "'break'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "ID", "STR", "INT", "COMMENT", "WS",
	}
	staticData.ruleNames = []string{
		"program", "expression", "operationOu", "operationEt", "operationComparaison",
		"operationAddition", "operationMultiplication", "expressionUnaire",
		"sequenceInstruction", "operationNegation", "expressionValeur", "expressionValeur2",
		"operationSi", "operationTantque", "operationBoucle", "definition",
		"declaration", "declarationType", "declarationType2", "declarationChamp",
		"declarationFonction", "declarationValeur", "constantes", "identifiant",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 45, 303, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 3, 1, 55, 8, 1, 1, 2, 1, 2, 1, 2, 5, 2, 60, 8, 2, 10, 2, 12, 2, 63,
		9, 2, 1, 3, 1, 3, 1, 3, 5, 3, 68, 8, 3, 10, 3, 12, 3, 71, 9, 3, 1, 4, 1,
		4, 1, 4, 3, 4, 76, 8, 4, 1, 5, 1, 5, 1, 5, 5, 5, 81, 8, 5, 10, 5, 12, 5,
		84, 9, 5, 1, 6, 1, 6, 1, 6, 5, 6, 89, 8, 6, 10, 6, 12, 6, 92, 9, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 102, 8, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 5, 8, 108, 8, 8, 10, 8, 12, 8, 111, 9, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 3, 10, 120, 8, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 5, 11, 126, 8, 11, 10, 11, 12, 11, 129, 9, 11, 3, 11, 131, 8, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 5, 11, 145, 8, 11, 10, 11, 12, 11, 148, 9, 11, 3, 11, 150, 8, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 160, 8,
		11, 10, 11, 12, 11, 163, 9, 11, 3, 11, 165, 8, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 176, 8, 11, 10, 11, 12,
		11, 179, 9, 11, 3, 11, 181, 8, 11, 1, 11, 3, 11, 184, 8, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		3, 12, 198, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 4, 15, 216,
		8, 15, 11, 15, 12, 15, 217, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 224, 8,
		15, 10, 15, 12, 15, 227, 9, 15, 3, 15, 229, 8, 15, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 16, 3, 16, 236, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 251, 8, 18,
		10, 18, 12, 18, 254, 9, 18, 3, 18, 256, 8, 18, 1, 18, 3, 18, 259, 8, 18,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 5,
		20, 271, 8, 20, 10, 20, 12, 20, 274, 9, 20, 3, 20, 276, 8, 20, 1, 20, 1,
		20, 1, 20, 3, 20, 281, 8, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21,
		1, 21, 3, 21, 290, 8, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1,
		22, 3, 22, 299, 8, 22, 1, 23, 1, 23, 1, 23, 0, 0, 24, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
		0, 3, 1, 0, 4, 9, 1, 0, 10, 11, 1, 0, 12, 13, 322, 0, 48, 1, 0, 0, 0, 2,
		51, 1, 0, 0, 0, 4, 56, 1, 0, 0, 0, 6, 64, 1, 0, 0, 0, 8, 72, 1, 0, 0, 0,
		10, 77, 1, 0, 0, 0, 12, 85, 1, 0, 0, 0, 14, 101, 1, 0, 0, 0, 16, 103, 1,
		0, 0, 0, 18, 114, 1, 0, 0, 0, 20, 117, 1, 0, 0, 0, 22, 183, 1, 0, 0, 0,
		24, 197, 1, 0, 0, 0, 26, 199, 1, 0, 0, 0, 28, 204, 1, 0, 0, 0, 30, 213,
		1, 0, 0, 0, 32, 235, 1, 0, 0, 0, 34, 237, 1, 0, 0, 0, 36, 258, 1, 0, 0,
		0, 38, 260, 1, 0, 0, 0, 40, 264, 1, 0, 0, 0, 42, 285, 1, 0, 0, 0, 44, 298,
		1, 0, 0, 0, 46, 300, 1, 0, 0, 0, 48, 49, 3, 2, 1, 0, 49, 50, 5, 0, 0, 1,
		50, 1, 1, 0, 0, 0, 51, 54, 3, 4, 2, 0, 52, 53, 5, 1, 0, 0, 53, 55, 3, 4,
		2, 0, 54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 3, 1, 0, 0, 0, 56, 61,
		3, 6, 3, 0, 57, 58, 5, 2, 0, 0, 58, 60, 3, 6, 3, 0, 59, 57, 1, 0, 0, 0,
		60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 5, 1, 0,
		0, 0, 63, 61, 1, 0, 0, 0, 64, 69, 3, 8, 4, 0, 65, 66, 5, 3, 0, 0, 66, 68,
		3, 8, 4, 0, 67, 65, 1, 0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0,
		69, 70, 1, 0, 0, 0, 70, 7, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 75, 3, 10,
		5, 0, 73, 74, 7, 0, 0, 0, 74, 76, 3, 10, 5, 0, 75, 73, 1, 0, 0, 0, 75,
		76, 1, 0, 0, 0, 76, 9, 1, 0, 0, 0, 77, 82, 3, 12, 6, 0, 78, 79, 7, 1, 0,
		0, 79, 81, 3, 12, 6, 0, 80, 78, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80,
		1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 11, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0,
		85, 90, 3, 14, 7, 0, 86, 87, 7, 2, 0, 0, 87, 89, 3, 14, 7, 0, 88, 86, 1,
		0, 0, 0, 89, 92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91,
		13, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 93, 102, 3, 16, 8, 0, 94, 102, 3, 18,
		9, 0, 95, 102, 3, 20, 10, 0, 96, 102, 3, 24, 12, 0, 97, 102, 3, 26, 13,
		0, 98, 102, 3, 28, 14, 0, 99, 102, 3, 30, 15, 0, 100, 102, 3, 44, 22, 0,
		101, 93, 1, 0, 0, 0, 101, 94, 1, 0, 0, 0, 101, 95, 1, 0, 0, 0, 101, 96,
		1, 0, 0, 0, 101, 97, 1, 0, 0, 0, 101, 98, 1, 0, 0, 0, 101, 99, 1, 0, 0,
		0, 101, 100, 1, 0, 0, 0, 102, 15, 1, 0, 0, 0, 103, 104, 5, 14, 0, 0, 104,
		109, 3, 2, 1, 0, 105, 106, 5, 15, 0, 0, 106, 108, 3, 2, 1, 0, 107, 105,
		1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0,
		0, 0, 110, 112, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112, 113, 5, 16, 0, 0,
		113, 17, 1, 0, 0, 0, 114, 115, 5, 11, 0, 0, 115, 116, 3, 14, 7, 0, 116,
		19, 1, 0, 0, 0, 117, 119, 3, 46, 23, 0, 118, 120, 3, 22, 11, 0, 119, 118,
		1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 21, 1, 0, 0, 0, 121, 130, 5, 14,
		0, 0, 122, 127, 3, 2, 1, 0, 123, 124, 5, 17, 0, 0, 124, 126, 3, 2, 1, 0,
		125, 123, 1, 0, 0, 0, 126, 129, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127,
		128, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 130, 122,
		1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 184, 5, 16,
		0, 0, 133, 134, 5, 18, 0, 0, 134, 135, 3, 2, 1, 0, 135, 149, 5, 19, 0,
		0, 136, 137, 5, 20, 0, 0, 137, 150, 3, 14, 7, 0, 138, 139, 5, 21, 0, 0,
		139, 145, 3, 46, 23, 0, 140, 141, 5, 18, 0, 0, 141, 142, 3, 2, 1, 0, 142,
		143, 5, 19, 0, 0, 143, 145, 1, 0, 0, 0, 144, 138, 1, 0, 0, 0, 144, 140,
		1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0,
		0, 0, 147, 150, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 149, 136, 1, 0, 0, 0,
		149, 146, 1, 0, 0, 0, 150, 165, 1, 0, 0, 0, 151, 152, 5, 21, 0, 0, 152,
		161, 3, 46, 23, 0, 153, 154, 5, 21, 0, 0, 154, 160, 3, 46, 23, 0, 155,
		156, 5, 18, 0, 0, 156, 157, 3, 2, 1, 0, 157, 158, 5, 19, 0, 0, 158, 160,
		1, 0, 0, 0, 159, 153, 1, 0, 0, 0, 159, 155, 1, 0, 0, 0, 160, 163, 1, 0,
		0, 0, 161, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 165, 1, 0, 0, 0,
		163, 161, 1, 0, 0, 0, 164, 133, 1, 0, 0, 0, 164, 151, 1, 0, 0, 0, 165,
		184, 1, 0, 0, 0, 166, 180, 5, 22, 0, 0, 167, 168, 3, 46, 23, 0, 168, 169,
		5, 4, 0, 0, 169, 177, 3, 2, 1, 0, 170, 171, 5, 17, 0, 0, 171, 172, 3, 46,
		23, 0, 172, 173, 5, 4, 0, 0, 173, 174, 3, 2, 1, 0, 174, 176, 1, 0, 0, 0,
		175, 170, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177,
		178, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 180, 167,
		1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 184, 5, 23,
		0, 0, 183, 121, 1, 0, 0, 0, 183, 164, 1, 0, 0, 0, 183, 166, 1, 0, 0, 0,
		184, 23, 1, 0, 0, 0, 185, 186, 5, 24, 0, 0, 186, 187, 3, 2, 1, 0, 187,
		188, 5, 25, 0, 0, 188, 189, 3, 14, 7, 0, 189, 198, 1, 0, 0, 0, 190, 191,
		5, 24, 0, 0, 191, 192, 3, 2, 1, 0, 192, 193, 5, 25, 0, 0, 193, 194, 3,
		14, 7, 0, 194, 195, 5, 26, 0, 0, 195, 196, 3, 14, 7, 0, 196, 198, 1, 0,
		0, 0, 197, 185, 1, 0, 0, 0, 197, 190, 1, 0, 0, 0, 198, 25, 1, 0, 0, 0,
		199, 200, 5, 27, 0, 0, 200, 201, 3, 2, 1, 0, 201, 202, 5, 28, 0, 0, 202,
		203, 3, 14, 7, 0, 203, 27, 1, 0, 0, 0, 204, 205, 5, 29, 0, 0, 205, 206,
		3, 46, 23, 0, 206, 207, 5, 1, 0, 0, 207, 208, 3, 2, 1, 0, 208, 209, 5,
		30, 0, 0, 209, 210, 3, 2, 1, 0, 210, 211, 5, 28, 0, 0, 211, 212, 3, 14,
		7, 0, 212, 29, 1, 0, 0, 0, 213, 215, 5, 31, 0, 0, 214, 216, 3, 32, 16,
		0, 215, 214, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 215, 1, 0, 0, 0, 217,
		218, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 228, 5, 32, 0, 0, 220, 225,
		3, 2, 1, 0, 221, 222, 5, 15, 0, 0, 222, 224, 3, 2, 1, 0, 223, 221, 1, 0,
		0, 0, 224, 227, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0,
		226, 229, 1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 228, 220, 1, 0, 0, 0, 228,
		229, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 231, 5, 33, 0, 0, 231, 31,
		1, 0, 0, 0, 232, 236, 3, 34, 17, 0, 233, 236, 3, 40, 20, 0, 234, 236, 3,
		42, 21, 0, 235, 232, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 235, 234, 1, 0,
		0, 0, 236, 33, 1, 0, 0, 0, 237, 238, 5, 34, 0, 0, 238, 239, 3, 46, 23,
		0, 239, 240, 5, 4, 0, 0, 240, 241, 3, 36, 18, 0, 241, 35, 1, 0, 0, 0, 242,
		259, 3, 46, 23, 0, 243, 244, 5, 35, 0, 0, 244, 245, 5, 20, 0, 0, 245, 259,
		3, 46, 23, 0, 246, 255, 5, 22, 0, 0, 247, 252, 3, 38, 19, 0, 248, 249,
		5, 17, 0, 0, 249, 251, 3, 38, 19, 0, 250, 248, 1, 0, 0, 0, 251, 254, 1,
		0, 0, 0, 252, 250, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 256, 1, 0, 0,
		0, 254, 252, 1, 0, 0, 0, 255, 247, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256,
		257, 1, 0, 0, 0, 257, 259, 5, 23, 0, 0, 258, 242, 1, 0, 0, 0, 258, 243,
		1, 0, 0, 0, 258, 246, 1, 0, 0, 0, 259, 37, 1, 0, 0, 0, 260, 261, 3, 46,
		23, 0, 261, 262, 5, 36, 0, 0, 262, 263, 3, 46, 23, 0, 263, 39, 1, 0, 0,
		0, 264, 265, 5, 37, 0, 0, 265, 266, 3, 46, 23, 0, 266, 275, 5, 14, 0, 0,
		267, 272, 3, 38, 19, 0, 268, 269, 5, 17, 0, 0, 269, 271, 3, 38, 19, 0,
		270, 268, 1, 0, 0, 0, 271, 274, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 272,
		273, 1, 0, 0, 0, 273, 276, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 275, 267,
		1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 280, 5, 16,
		0, 0, 278, 279, 5, 36, 0, 0, 279, 281, 3, 46, 23, 0, 280, 278, 1, 0, 0,
		0, 280, 281, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 283, 5, 4, 0, 0, 283,
		284, 3, 2, 1, 0, 284, 41, 1, 0, 0, 0, 285, 286, 5, 38, 0, 0, 286, 289,
		3, 46, 23, 0, 287, 288, 5, 36, 0, 0, 288, 290, 3, 46, 23, 0, 289, 287,
		1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 292, 5, 1,
		0, 0, 292, 293, 3, 2, 1, 0, 293, 43, 1, 0, 0, 0, 294, 299, 5, 42, 0, 0,
		295, 299, 5, 43, 0, 0, 296, 299, 5, 39, 0, 0, 297, 299, 5, 40, 0, 0, 298,
		294, 1, 0, 0, 0, 298, 295, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 298, 297,
		1, 0, 0, 0, 299, 45, 1, 0, 0, 0, 300, 301, 5, 41, 0, 0, 301, 47, 1, 0,
		0, 0, 33, 54, 61, 69, 75, 82, 90, 101, 109, 119, 127, 130, 144, 146, 149,
		159, 161, 164, 177, 180, 183, 197, 217, 225, 228, 235, 252, 255, 258, 272,
		275, 280, 289, 298,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// tigerParserInit initializes any static state used to implement tigerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewtigerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TigerParserInit() {
	staticData := &tigerParserStaticData
	staticData.once.Do(tigerParserInit)
}

// NewtigerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewtigerParser(input antlr.TokenStream) *tigerParser {
	TigerParserInit()
	this := new(tigerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &tigerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "tiger.g4"

	return this
}

// tigerParser tokens.
const (
	tigerParserEOF     = antlr.TokenEOF
	tigerParserT__0    = 1
	tigerParserT__1    = 2
	tigerParserT__2    = 3
	tigerParserT__3    = 4
	tigerParserT__4    = 5
	tigerParserT__5    = 6
	tigerParserT__6    = 7
	tigerParserT__7    = 8
	tigerParserT__8    = 9
	tigerParserT__9    = 10
	tigerParserT__10   = 11
	tigerParserT__11   = 12
	tigerParserT__12   = 13
	tigerParserT__13   = 14
	tigerParserT__14   = 15
	tigerParserT__15   = 16
	tigerParserT__16   = 17
	tigerParserT__17   = 18
	tigerParserT__18   = 19
	tigerParserT__19   = 20
	tigerParserT__20   = 21
	tigerParserT__21   = 22
	tigerParserT__22   = 23
	tigerParserT__23   = 24
	tigerParserT__24   = 25
	tigerParserT__25   = 26
	tigerParserT__26   = 27
	tigerParserT__27   = 28
	tigerParserT__28   = 29
	tigerParserT__29   = 30
	tigerParserT__30   = 31
	tigerParserT__31   = 32
	tigerParserT__32   = 33
	tigerParserT__33   = 34
	tigerParserT__34   = 35
	tigerParserT__35   = 36
	tigerParserT__36   = 37
	tigerParserT__37   = 38
	tigerParserT__38   = 39
	tigerParserT__39   = 40
	tigerParserID      = 41
	tigerParserSTR     = 42
	tigerParserINT     = 43
	tigerParserCOMMENT = 44
	tigerParserWS      = 45
)

// tigerParser rules.
const (
	tigerParserRULE_program                 = 0
	tigerParserRULE_expression              = 1
	tigerParserRULE_operationOu             = 2
	tigerParserRULE_operationEt             = 3
	tigerParserRULE_operationComparaison    = 4
	tigerParserRULE_operationAddition       = 5
	tigerParserRULE_operationMultiplication = 6
	tigerParserRULE_expressionUnaire        = 7
	tigerParserRULE_sequenceInstruction     = 8
	tigerParserRULE_operationNegation       = 9
	tigerParserRULE_expressionValeur        = 10
	tigerParserRULE_expressionValeur2       = 11
	tigerParserRULE_operationSi             = 12
	tigerParserRULE_operationTantque        = 13
	tigerParserRULE_operationBoucle         = 14
	tigerParserRULE_definition              = 15
	tigerParserRULE_declaration             = 16
	tigerParserRULE_declarationType         = 17
	tigerParserRULE_declarationType2        = 18
	tigerParserRULE_declarationChamp        = 19
	tigerParserRULE_declarationFonction     = 20
	tigerParserRULE_declarationValeur       = 21
	tigerParserRULE_constantes              = 22
	tigerParserRULE_identifiant             = 23
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	EOF() antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(tigerParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, tigerParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Expression()
	}
	{
		p.SetState(49)
		p.Match(tigerParserEOF)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOperationOu() []IOperationOuContext
	OperationOu(i int) IOperationOuContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllOperationOu() []IOperationOuContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperationOuContext); ok {
			len++
		}
	}

	tst := make([]IOperationOuContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperationOuContext); ok {
			tst[i] = t.(IOperationOuContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) OperationOu(i int) IOperationOuContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationOuContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperationOuContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, tigerParserRULE_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.OperationOu()
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tigerParserT__0 {
		{
			p.SetState(52)
			p.Match(tigerParserT__0)
		}
		{
			p.SetState(53)
			p.OperationOu()
		}

	}

	return localctx
}

// IOperationOuContext is an interface to support dynamic dispatch.
type IOperationOuContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOperationEt() []IOperationEtContext
	OperationEt(i int) IOperationEtContext

	// IsOperationOuContext differentiates from other interfaces.
	IsOperationOuContext()
}

type OperationOuContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationOuContext() *OperationOuContext {
	var p = new(OperationOuContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_operationOu
	return p
}

func (*OperationOuContext) IsOperationOuContext() {}

func NewOperationOuContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationOuContext {
	var p = new(OperationOuContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_operationOu

	return p
}

func (s *OperationOuContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationOuContext) AllOperationEt() []IOperationEtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperationEtContext); ok {
			len++
		}
	}

	tst := make([]IOperationEtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperationEtContext); ok {
			tst[i] = t.(IOperationEtContext)
			i++
		}
	}

	return tst
}

func (s *OperationOuContext) OperationEt(i int) IOperationEtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationEtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperationEtContext)
}

func (s *OperationOuContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationOuContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationOuContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitOperationOu(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) OperationOu() (localctx IOperationOuContext) {
	this := p
	_ = this

	localctx = NewOperationOuContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, tigerParserRULE_operationOu)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.OperationEt()
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tigerParserT__1 {
		{
			p.SetState(57)
			p.Match(tigerParserT__1)
		}
		{
			p.SetState(58)
			p.OperationEt()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOperationEtContext is an interface to support dynamic dispatch.
type IOperationEtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOperationComparaison() []IOperationComparaisonContext
	OperationComparaison(i int) IOperationComparaisonContext

	// IsOperationEtContext differentiates from other interfaces.
	IsOperationEtContext()
}

type OperationEtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationEtContext() *OperationEtContext {
	var p = new(OperationEtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_operationEt
	return p
}

func (*OperationEtContext) IsOperationEtContext() {}

func NewOperationEtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationEtContext {
	var p = new(OperationEtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_operationEt

	return p
}

func (s *OperationEtContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationEtContext) AllOperationComparaison() []IOperationComparaisonContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperationComparaisonContext); ok {
			len++
		}
	}

	tst := make([]IOperationComparaisonContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperationComparaisonContext); ok {
			tst[i] = t.(IOperationComparaisonContext)
			i++
		}
	}

	return tst
}

func (s *OperationEtContext) OperationComparaison(i int) IOperationComparaisonContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationComparaisonContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperationComparaisonContext)
}

func (s *OperationEtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationEtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationEtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitOperationEt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) OperationEt() (localctx IOperationEtContext) {
	this := p
	_ = this

	localctx = NewOperationEtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, tigerParserRULE_operationEt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.OperationComparaison()
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tigerParserT__2 {
		{
			p.SetState(65)
			p.Match(tigerParserT__2)
		}
		{
			p.SetState(66)
			p.OperationComparaison()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOperationComparaisonContext is an interface to support dynamic dispatch.
type IOperationComparaisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOperationAddition() []IOperationAdditionContext
	OperationAddition(i int) IOperationAdditionContext

	// IsOperationComparaisonContext differentiates from other interfaces.
	IsOperationComparaisonContext()
}

type OperationComparaisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationComparaisonContext() *OperationComparaisonContext {
	var p = new(OperationComparaisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_operationComparaison
	return p
}

func (*OperationComparaisonContext) IsOperationComparaisonContext() {}

func NewOperationComparaisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationComparaisonContext {
	var p = new(OperationComparaisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_operationComparaison

	return p
}

func (s *OperationComparaisonContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationComparaisonContext) AllOperationAddition() []IOperationAdditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperationAdditionContext); ok {
			len++
		}
	}

	tst := make([]IOperationAdditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperationAdditionContext); ok {
			tst[i] = t.(IOperationAdditionContext)
			i++
		}
	}

	return tst
}

func (s *OperationComparaisonContext) OperationAddition(i int) IOperationAdditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationAdditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperationAdditionContext)
}

func (s *OperationComparaisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationComparaisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationComparaisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitOperationComparaison(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) OperationComparaison() (localctx IOperationComparaisonContext) {
	this := p
	_ = this

	localctx = NewOperationComparaisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, tigerParserRULE_operationComparaison)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.OperationAddition()
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1008) != 0 {
		{
			p.SetState(73)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1008) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(74)
			p.OperationAddition()
		}

	}

	return localctx
}

// IOperationAdditionContext is an interface to support dynamic dispatch.
type IOperationAdditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOperationMultiplication() []IOperationMultiplicationContext
	OperationMultiplication(i int) IOperationMultiplicationContext

	// IsOperationAdditionContext differentiates from other interfaces.
	IsOperationAdditionContext()
}

type OperationAdditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationAdditionContext() *OperationAdditionContext {
	var p = new(OperationAdditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_operationAddition
	return p
}

func (*OperationAdditionContext) IsOperationAdditionContext() {}

func NewOperationAdditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationAdditionContext {
	var p = new(OperationAdditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_operationAddition

	return p
}

func (s *OperationAdditionContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationAdditionContext) AllOperationMultiplication() []IOperationMultiplicationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperationMultiplicationContext); ok {
			len++
		}
	}

	tst := make([]IOperationMultiplicationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperationMultiplicationContext); ok {
			tst[i] = t.(IOperationMultiplicationContext)
			i++
		}
	}

	return tst
}

func (s *OperationAdditionContext) OperationMultiplication(i int) IOperationMultiplicationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationMultiplicationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperationMultiplicationContext)
}

func (s *OperationAdditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationAdditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationAdditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitOperationAddition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) OperationAddition() (localctx IOperationAdditionContext) {
	this := p
	_ = this

	localctx = NewOperationAdditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, tigerParserRULE_operationAddition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.OperationMultiplication()
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tigerParserT__9 || _la == tigerParserT__10 {
		{
			p.SetState(78)
			_la = p.GetTokenStream().LA(1)

			if !(_la == tigerParserT__9 || _la == tigerParserT__10) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(79)
			p.OperationMultiplication()
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOperationMultiplicationContext is an interface to support dynamic dispatch.
type IOperationMultiplicationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpressionUnaire() []IExpressionUnaireContext
	ExpressionUnaire(i int) IExpressionUnaireContext

	// IsOperationMultiplicationContext differentiates from other interfaces.
	IsOperationMultiplicationContext()
}

type OperationMultiplicationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationMultiplicationContext() *OperationMultiplicationContext {
	var p = new(OperationMultiplicationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_operationMultiplication
	return p
}

func (*OperationMultiplicationContext) IsOperationMultiplicationContext() {}

func NewOperationMultiplicationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationMultiplicationContext {
	var p = new(OperationMultiplicationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_operationMultiplication

	return p
}

func (s *OperationMultiplicationContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationMultiplicationContext) AllExpressionUnaire() []IExpressionUnaireContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionUnaireContext); ok {
			len++
		}
	}

	tst := make([]IExpressionUnaireContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionUnaireContext); ok {
			tst[i] = t.(IExpressionUnaireContext)
			i++
		}
	}

	return tst
}

func (s *OperationMultiplicationContext) ExpressionUnaire(i int) IExpressionUnaireContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionUnaireContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionUnaireContext)
}

func (s *OperationMultiplicationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationMultiplicationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationMultiplicationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitOperationMultiplication(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) OperationMultiplication() (localctx IOperationMultiplicationContext) {
	this := p
	_ = this

	localctx = NewOperationMultiplicationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, tigerParserRULE_operationMultiplication)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.ExpressionUnaire()
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tigerParserT__11 || _la == tigerParserT__12 {
		{
			p.SetState(86)
			_la = p.GetTokenStream().LA(1)

			if !(_la == tigerParserT__11 || _la == tigerParserT__12) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(87)
			p.ExpressionUnaire()
		}

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressionUnaireContext is an interface to support dynamic dispatch.
type IExpressionUnaireContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SequenceInstruction() ISequenceInstructionContext
	OperationNegation() IOperationNegationContext
	ExpressionValeur() IExpressionValeurContext
	OperationSi() IOperationSiContext
	OperationTantque() IOperationTantqueContext
	OperationBoucle() IOperationBoucleContext
	Definition() IDefinitionContext
	Constantes() IConstantesContext

	// IsExpressionUnaireContext differentiates from other interfaces.
	IsExpressionUnaireContext()
}

type ExpressionUnaireContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionUnaireContext() *ExpressionUnaireContext {
	var p = new(ExpressionUnaireContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_expressionUnaire
	return p
}

func (*ExpressionUnaireContext) IsExpressionUnaireContext() {}

func NewExpressionUnaireContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionUnaireContext {
	var p = new(ExpressionUnaireContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_expressionUnaire

	return p
}

func (s *ExpressionUnaireContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionUnaireContext) SequenceInstruction() ISequenceInstructionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequenceInstructionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISequenceInstructionContext)
}

func (s *ExpressionUnaireContext) OperationNegation() IOperationNegationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationNegationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperationNegationContext)
}

func (s *ExpressionUnaireContext) ExpressionValeur() IExpressionValeurContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionValeurContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionValeurContext)
}

func (s *ExpressionUnaireContext) OperationSi() IOperationSiContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationSiContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperationSiContext)
}

func (s *ExpressionUnaireContext) OperationTantque() IOperationTantqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationTantqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperationTantqueContext)
}

func (s *ExpressionUnaireContext) OperationBoucle() IOperationBoucleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationBoucleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperationBoucleContext)
}

func (s *ExpressionUnaireContext) Definition() IDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *ExpressionUnaireContext) Constantes() IConstantesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantesContext)
}

func (s *ExpressionUnaireContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionUnaireContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionUnaireContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitExpressionUnaire(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) ExpressionUnaire() (localctx IExpressionUnaireContext) {
	this := p
	_ = this

	localctx = NewExpressionUnaireContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, tigerParserRULE_expressionUnaire)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tigerParserT__13:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.SequenceInstruction()
		}

	case tigerParserT__10:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.OperationNegation()
		}

	case tigerParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(95)
			p.ExpressionValeur()
		}

	case tigerParserT__23:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(96)
			p.OperationSi()
		}

	case tigerParserT__26:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(97)
			p.OperationTantque()
		}

	case tigerParserT__28:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(98)
			p.OperationBoucle()
		}

	case tigerParserT__30:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(99)
			p.Definition()
		}

	case tigerParserT__38, tigerParserT__39, tigerParserSTR, tigerParserINT:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(100)
			p.Constantes()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISequenceInstructionContext is an interface to support dynamic dispatch.
type ISequenceInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSequenceInstructionContext differentiates from other interfaces.
	IsSequenceInstructionContext()
}

type SequenceInstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequenceInstructionContext() *SequenceInstructionContext {
	var p = new(SequenceInstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_sequenceInstruction
	return p
}

func (*SequenceInstructionContext) IsSequenceInstructionContext() {}

func NewSequenceInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SequenceInstructionContext {
	var p = new(SequenceInstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_sequenceInstruction

	return p
}

func (s *SequenceInstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *SequenceInstructionContext) CopyFrom(ctx *SequenceInstructionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SequenceInstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequenceInstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SequenceContext struct {
	*SequenceInstructionContext
}

func NewSequenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SequenceContext {
	var p = new(SequenceContext)

	p.SequenceInstructionContext = NewEmptySequenceInstructionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SequenceInstructionContext))

	return p
}

func (s *SequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequenceContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *SequenceContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SequenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitSequence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) SequenceInstruction() (localctx ISequenceInstructionContext) {
	this := p
	_ = this

	localctx = NewSequenceInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, tigerParserRULE_sequenceInstruction)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewSequenceContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(tigerParserT__13)
	}
	{
		p.SetState(104)
		p.Expression()
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tigerParserT__14 {
		{
			p.SetState(105)
			p.Match(tigerParserT__14)
		}
		{
			p.SetState(106)
			p.Expression()
		}

		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(112)
		p.Match(tigerParserT__15)
	}

	return localctx
}

// IOperationNegationContext is an interface to support dynamic dispatch.
type IOperationNegationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperationNegationContext differentiates from other interfaces.
	IsOperationNegationContext()
}

type OperationNegationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationNegationContext() *OperationNegationContext {
	var p = new(OperationNegationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_operationNegation
	return p
}

func (*OperationNegationContext) IsOperationNegationContext() {}

func NewOperationNegationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationNegationContext {
	var p = new(OperationNegationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_operationNegation

	return p
}

func (s *OperationNegationContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationNegationContext) CopyFrom(ctx *OperationNegationContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *OperationNegationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationNegationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NegationContext struct {
	*OperationNegationContext
}

func NewNegationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegationContext {
	var p = new(NegationContext)

	p.OperationNegationContext = NewEmptyOperationNegationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationNegationContext))

	return p
}

func (s *NegationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegationContext) ExpressionUnaire() IExpressionUnaireContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionUnaireContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionUnaireContext)
}

func (s *NegationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitNegation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) OperationNegation() (localctx IOperationNegationContext) {
	this := p
	_ = this

	localctx = NewOperationNegationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, tigerParserRULE_operationNegation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewNegationContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(tigerParserT__10)
	}
	{
		p.SetState(115)
		p.ExpressionUnaire()
	}

	return localctx
}

// IExpressionValeurContext is an interface to support dynamic dispatch.
type IExpressionValeurContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionValeurContext differentiates from other interfaces.
	IsExpressionValeurContext()
}

type ExpressionValeurContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionValeurContext() *ExpressionValeurContext {
	var p = new(ExpressionValeurContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_expressionValeur
	return p
}

func (*ExpressionValeurContext) IsExpressionValeurContext() {}

func NewExpressionValeurContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionValeurContext {
	var p = new(ExpressionValeurContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_expressionValeur

	return p
}

func (s *ExpressionValeurContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionValeurContext) CopyFrom(ctx *ExpressionValeurContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionValeurContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionValeurContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionIdentifiantContext struct {
	*ExpressionValeurContext
}

func NewExpressionIdentifiantContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionIdentifiantContext {
	var p = new(ExpressionIdentifiantContext)

	p.ExpressionValeurContext = NewEmptyExpressionValeurContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionValeurContext))

	return p
}

func (s *ExpressionIdentifiantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionIdentifiantContext) Identifiant() IIdentifiantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifiantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifiantContext)
}

func (s *ExpressionIdentifiantContext) ExpressionValeur2() IExpressionValeur2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionValeur2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionValeur2Context)
}

func (s *ExpressionIdentifiantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitExpressionIdentifiant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) ExpressionValeur() (localctx IExpressionValeurContext) {
	this := p
	_ = this

	localctx = NewExpressionValeurContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, tigerParserRULE_expressionValeur)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewExpressionIdentifiantContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Identifiant()
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6569984) != 0 {
		{
			p.SetState(118)
			p.ExpressionValeur2()
		}

	}

	return localctx
}

// IExpressionValeur2Context is an interface to support dynamic dispatch.
type IExpressionValeur2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionValeur2Context differentiates from other interfaces.
	IsExpressionValeur2Context()
}

type ExpressionValeur2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionValeur2Context() *ExpressionValeur2Context {
	var p = new(ExpressionValeur2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_expressionValeur2
	return p
}

func (*ExpressionValeur2Context) IsExpressionValeur2Context() {}

func NewExpressionValeur2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionValeur2Context {
	var p = new(ExpressionValeur2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_expressionValeur2

	return p
}

func (s *ExpressionValeur2Context) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionValeur2Context) CopyFrom(ctx *ExpressionValeur2Context) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionValeur2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionValeur2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AppelFonctionContext struct {
	*ExpressionValeur2Context
}

func NewAppelFonctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AppelFonctionContext {
	var p = new(AppelFonctionContext)

	p.ExpressionValeur2Context = NewEmptyExpressionValeur2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionValeur2Context))

	return p
}

func (s *AppelFonctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppelFonctionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AppelFonctionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AppelFonctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitAppelFonction(s)

	default:
		return t.VisitChildren(s)
	}
}

type InstanciationTypeContext struct {
	*ExpressionValeur2Context
}

func NewInstanciationTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InstanciationTypeContext {
	var p = new(InstanciationTypeContext)

	p.ExpressionValeur2Context = NewEmptyExpressionValeur2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionValeur2Context))

	return p
}

func (s *InstanciationTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanciationTypeContext) AllIdentifiant() []IIdentifiantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifiantContext); ok {
			len++
		}
	}

	tst := make([]IIdentifiantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifiantContext); ok {
			tst[i] = t.(IIdentifiantContext)
			i++
		}
	}

	return tst
}

func (s *InstanciationTypeContext) Identifiant(i int) IIdentifiantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifiantContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifiantContext)
}

func (s *InstanciationTypeContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *InstanciationTypeContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InstanciationTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitInstanciationType(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListeAccesContext struct {
	*ExpressionValeur2Context
}

func NewListeAccesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListeAccesContext {
	var p = new(ListeAccesContext)

	p.ExpressionValeur2Context = NewEmptyExpressionValeur2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionValeur2Context))

	return p
}

func (s *ListeAccesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListeAccesContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ListeAccesContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListeAccesContext) AllIdentifiant() []IIdentifiantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifiantContext); ok {
			len++
		}
	}

	tst := make([]IIdentifiantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifiantContext); ok {
			tst[i] = t.(IIdentifiantContext)
			i++
		}
	}

	return tst
}

func (s *ListeAccesContext) Identifiant(i int) IIdentifiantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifiantContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifiantContext)
}

func (s *ListeAccesContext) ExpressionUnaire() IExpressionUnaireContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionUnaireContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionUnaireContext)
}

func (s *ListeAccesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitListeAcces(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) ExpressionValeur2() (localctx IExpressionValeur2Context) {
	this := p
	_ = this

	localctx = NewExpressionValeur2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, tigerParserRULE_expressionValeur2)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(183)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tigerParserT__13:
		localctx = NewAppelFonctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.Match(tigerParserT__13)
		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17045265598464) != 0 {
			{
				p.SetState(122)
				p.Expression()
			}
			p.SetState(127)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == tigerParserT__16 {
				{
					p.SetState(123)
					p.Match(tigerParserT__16)
				}
				{
					p.SetState(124)
					p.Expression()
				}

				p.SetState(129)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(132)
			p.Match(tigerParserT__15)
		}

	case tigerParserT__17, tigerParserT__20:
		localctx = NewListeAccesContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(164)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case tigerParserT__17:
			{
				p.SetState(133)
				p.Match(tigerParserT__17)
			}
			{
				p.SetState(134)
				p.Expression()
			}
			{
				p.SetState(135)
				p.Match(tigerParserT__18)
			}
			p.SetState(149)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case tigerParserT__19:
				{
					p.SetState(136)
					p.Match(tigerParserT__19)
				}
				{
					p.SetState(137)
					p.ExpressionUnaire()
				}

			case tigerParserEOF, tigerParserT__0, tigerParserT__1, tigerParserT__2, tigerParserT__3, tigerParserT__4, tigerParserT__5, tigerParserT__6, tigerParserT__7, tigerParserT__8, tigerParserT__9, tigerParserT__10, tigerParserT__11, tigerParserT__12, tigerParserT__14, tigerParserT__15, tigerParserT__16, tigerParserT__17, tigerParserT__18, tigerParserT__20, tigerParserT__22, tigerParserT__24, tigerParserT__25, tigerParserT__27, tigerParserT__29, tigerParserT__31, tigerParserT__32, tigerParserT__33, tigerParserT__36, tigerParserT__37:
				p.SetState(146)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == tigerParserT__17 || _la == tigerParserT__20 {
					p.SetState(144)
					p.GetErrorHandler().Sync(p)

					switch p.GetTokenStream().LA(1) {
					case tigerParserT__20:
						{
							p.SetState(138)
							p.Match(tigerParserT__20)
						}
						{
							p.SetState(139)
							p.Identifiant()
						}

					case tigerParserT__17:
						{
							p.SetState(140)
							p.Match(tigerParserT__17)
						}
						{
							p.SetState(141)
							p.Expression()
						}
						{
							p.SetState(142)
							p.Match(tigerParserT__18)
						}

					default:
						panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					}

					p.SetState(148)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		case tigerParserT__20:
			{
				p.SetState(151)
				p.Match(tigerParserT__20)
			}
			{
				p.SetState(152)
				p.Identifiant()
			}
			p.SetState(161)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == tigerParserT__17 || _la == tigerParserT__20 {
				p.SetState(159)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case tigerParserT__20:
					{
						p.SetState(153)
						p.Match(tigerParserT__20)
					}
					{
						p.SetState(154)
						p.Identifiant()
					}

				case tigerParserT__17:
					{
						p.SetState(155)
						p.Match(tigerParserT__17)
					}
					{
						p.SetState(156)
						p.Expression()
					}
					{
						p.SetState(157)
						p.Match(tigerParserT__18)
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(163)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	case tigerParserT__21:
		localctx = NewInstanciationTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(166)
			p.Match(tigerParserT__21)
		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == tigerParserID {
			{
				p.SetState(167)
				p.Identifiant()
			}
			{
				p.SetState(168)
				p.Match(tigerParserT__3)
			}
			{
				p.SetState(169)
				p.Expression()
			}
			p.SetState(177)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == tigerParserT__16 {
				{
					p.SetState(170)
					p.Match(tigerParserT__16)
				}
				{
					p.SetState(171)
					p.Identifiant()
				}
				{
					p.SetState(172)
					p.Match(tigerParserT__3)
				}
				{
					p.SetState(173)
					p.Expression()
				}

				p.SetState(179)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(182)
			p.Match(tigerParserT__22)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperationSiContext is an interface to support dynamic dispatch.
type IOperationSiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperationSiContext differentiates from other interfaces.
	IsOperationSiContext()
}

type OperationSiContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationSiContext() *OperationSiContext {
	var p = new(OperationSiContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_operationSi
	return p
}

func (*OperationSiContext) IsOperationSiContext() {}

func NewOperationSiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationSiContext {
	var p = new(OperationSiContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_operationSi

	return p
}

func (s *OperationSiContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationSiContext) CopyFrom(ctx *OperationSiContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *OperationSiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationSiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SiAlorsContext struct {
	*OperationSiContext
}

func NewSiAlorsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SiAlorsContext {
	var p = new(SiAlorsContext)

	p.OperationSiContext = NewEmptyOperationSiContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationSiContext))

	return p
}

func (s *SiAlorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SiAlorsContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SiAlorsContext) ExpressionUnaire() IExpressionUnaireContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionUnaireContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionUnaireContext)
}

func (s *SiAlorsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitSiAlors(s)

	default:
		return t.VisitChildren(s)
	}
}

type SiAlorsSinonContext struct {
	*OperationSiContext
}

func NewSiAlorsSinonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SiAlorsSinonContext {
	var p = new(SiAlorsSinonContext)

	p.OperationSiContext = NewEmptyOperationSiContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationSiContext))

	return p
}

func (s *SiAlorsSinonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SiAlorsSinonContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SiAlorsSinonContext) AllExpressionUnaire() []IExpressionUnaireContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionUnaireContext); ok {
			len++
		}
	}

	tst := make([]IExpressionUnaireContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionUnaireContext); ok {
			tst[i] = t.(IExpressionUnaireContext)
			i++
		}
	}

	return tst
}

func (s *SiAlorsSinonContext) ExpressionUnaire(i int) IExpressionUnaireContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionUnaireContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionUnaireContext)
}

func (s *SiAlorsSinonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitSiAlorsSinon(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) OperationSi() (localctx IOperationSiContext) {
	this := p
	_ = this

	localctx = NewOperationSiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, tigerParserRULE_operationSi)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSiAlorsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.Match(tigerParserT__23)
		}
		{
			p.SetState(186)
			p.Expression()
		}
		{
			p.SetState(187)
			p.Match(tigerParserT__24)
		}
		{
			p.SetState(188)
			p.ExpressionUnaire()
		}

	case 2:
		localctx = NewSiAlorsSinonContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Match(tigerParserT__23)
		}
		{
			p.SetState(191)
			p.Expression()
		}
		{
			p.SetState(192)
			p.Match(tigerParserT__24)
		}
		{
			p.SetState(193)
			p.ExpressionUnaire()
		}
		{
			p.SetState(194)
			p.Match(tigerParserT__25)
		}
		{
			p.SetState(195)
			p.ExpressionUnaire()
		}

	}

	return localctx
}

// IOperationTantqueContext is an interface to support dynamic dispatch.
type IOperationTantqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperationTantqueContext differentiates from other interfaces.
	IsOperationTantqueContext()
}

type OperationTantqueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationTantqueContext() *OperationTantqueContext {
	var p = new(OperationTantqueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_operationTantque
	return p
}

func (*OperationTantqueContext) IsOperationTantqueContext() {}

func NewOperationTantqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationTantqueContext {
	var p = new(OperationTantqueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_operationTantque

	return p
}

func (s *OperationTantqueContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationTantqueContext) CopyFrom(ctx *OperationTantqueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *OperationTantqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationTantqueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TantQueContext struct {
	*OperationTantqueContext
}

func NewTantQueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TantQueContext {
	var p = new(TantQueContext)

	p.OperationTantqueContext = NewEmptyOperationTantqueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationTantqueContext))

	return p
}

func (s *TantQueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TantQueContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TantQueContext) ExpressionUnaire() IExpressionUnaireContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionUnaireContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionUnaireContext)
}

func (s *TantQueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitTantQue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) OperationTantque() (localctx IOperationTantqueContext) {
	this := p
	_ = this

	localctx = NewOperationTantqueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, tigerParserRULE_operationTantque)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewTantQueContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(tigerParserT__26)
	}
	{
		p.SetState(200)
		p.Expression()
	}
	{
		p.SetState(201)
		p.Match(tigerParserT__27)
	}
	{
		p.SetState(202)
		p.ExpressionUnaire()
	}

	return localctx
}

// IOperationBoucleContext is an interface to support dynamic dispatch.
type IOperationBoucleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperationBoucleContext differentiates from other interfaces.
	IsOperationBoucleContext()
}

type OperationBoucleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationBoucleContext() *OperationBoucleContext {
	var p = new(OperationBoucleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_operationBoucle
	return p
}

func (*OperationBoucleContext) IsOperationBoucleContext() {}

func NewOperationBoucleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationBoucleContext {
	var p = new(OperationBoucleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_operationBoucle

	return p
}

func (s *OperationBoucleContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationBoucleContext) CopyFrom(ctx *OperationBoucleContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *OperationBoucleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationBoucleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PourContext struct {
	*OperationBoucleContext
}

func NewPourContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PourContext {
	var p = new(PourContext)

	p.OperationBoucleContext = NewEmptyOperationBoucleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationBoucleContext))

	return p
}

func (s *PourContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PourContext) Identifiant() IIdentifiantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifiantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifiantContext)
}

func (s *PourContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PourContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PourContext) ExpressionUnaire() IExpressionUnaireContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionUnaireContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionUnaireContext)
}

func (s *PourContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitPour(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) OperationBoucle() (localctx IOperationBoucleContext) {
	this := p
	_ = this

	localctx = NewOperationBoucleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, tigerParserRULE_operationBoucle)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewPourContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Match(tigerParserT__28)
	}
	{
		p.SetState(205)
		p.Identifiant()
	}
	{
		p.SetState(206)
		p.Match(tigerParserT__0)
	}
	{
		p.SetState(207)
		p.Expression()
	}
	{
		p.SetState(208)
		p.Match(tigerParserT__29)
	}
	{
		p.SetState(209)
		p.Expression()
	}
	{
		p.SetState(210)
		p.Match(tigerParserT__27)
	}
	{
		p.SetState(211)
		p.ExpressionUnaire()
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDeclaration() []IDeclarationContext
	Declaration(i int) IDeclarationContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *DefinitionContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *DefinitionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *DefinitionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) Definition() (localctx IDefinitionContext) {
	this := p
	_ = this

	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, tigerParserRULE_definition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Match(tigerParserT__30)
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&429496729600) != 0) {
		{
			p.SetState(214)
			p.Declaration()
		}

		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(219)
		p.Match(tigerParserT__31)
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17045265598464) != 0 {
		{
			p.SetState(220)
			p.Expression()
		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == tigerParserT__14 {
			{
				p.SetState(221)
				p.Match(tigerParserT__14)
			}
			{
				p.SetState(222)
				p.Expression()
			}

			p.SetState(227)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(230)
		p.Match(tigerParserT__32)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DeclarationType() IDeclarationTypeContext
	DeclarationFonction() IDeclarationFonctionContext
	DeclarationValeur() IDeclarationValeurContext

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) DeclarationType() IDeclarationTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationTypeContext)
}

func (s *DeclarationContext) DeclarationFonction() IDeclarationFonctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationFonctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationFonctionContext)
}

func (s *DeclarationContext) DeclarationValeur() IDeclarationValeurContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationValeurContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationValeurContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) Declaration() (localctx IDeclarationContext) {
	this := p
	_ = this

	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, tigerParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(235)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tigerParserT__33:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)
			p.DeclarationType()
		}

	case tigerParserT__36:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(233)
			p.DeclarationFonction()
		}

	case tigerParserT__37:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(234)
			p.DeclarationValeur()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDeclarationTypeContext is an interface to support dynamic dispatch.
type IDeclarationTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifiant() IIdentifiantContext
	DeclarationType2() IDeclarationType2Context

	// IsDeclarationTypeContext differentiates from other interfaces.
	IsDeclarationTypeContext()
}

type DeclarationTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationTypeContext() *DeclarationTypeContext {
	var p = new(DeclarationTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_declarationType
	return p
}

func (*DeclarationTypeContext) IsDeclarationTypeContext() {}

func NewDeclarationTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationTypeContext {
	var p = new(DeclarationTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_declarationType

	return p
}

func (s *DeclarationTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationTypeContext) Identifiant() IIdentifiantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifiantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifiantContext)
}

func (s *DeclarationTypeContext) DeclarationType2() IDeclarationType2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationType2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationType2Context)
}

func (s *DeclarationTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitDeclarationType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) DeclarationType() (localctx IDeclarationTypeContext) {
	this := p
	_ = this

	localctx = NewDeclarationTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, tigerParserRULE_declarationType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Match(tigerParserT__33)
	}
	{
		p.SetState(238)
		p.Identifiant()
	}
	{
		p.SetState(239)
		p.Match(tigerParserT__3)
	}
	{
		p.SetState(240)
		p.DeclarationType2()
	}

	return localctx
}

// IDeclarationType2Context is an interface to support dynamic dispatch.
type IDeclarationType2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclarationType2Context differentiates from other interfaces.
	IsDeclarationType2Context()
}

type DeclarationType2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationType2Context() *DeclarationType2Context {
	var p = new(DeclarationType2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_declarationType2
	return p
}

func (*DeclarationType2Context) IsDeclarationType2Context() {}

func NewDeclarationType2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationType2Context {
	var p = new(DeclarationType2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_declarationType2

	return p
}

func (s *DeclarationType2Context) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationType2Context) CopyFrom(ctx *DeclarationType2Context) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DeclarationType2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationType2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclarationRecordTypeContext struct {
	*DeclarationType2Context
}

func NewDeclarationRecordTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclarationRecordTypeContext {
	var p = new(DeclarationRecordTypeContext)

	p.DeclarationType2Context = NewEmptyDeclarationType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationType2Context))

	return p
}

func (s *DeclarationRecordTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationRecordTypeContext) AllDeclarationChamp() []IDeclarationChampContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationChampContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationChampContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationChampContext); ok {
			tst[i] = t.(IDeclarationChampContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationRecordTypeContext) DeclarationChamp(i int) IDeclarationChampContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationChampContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationChampContext)
}

func (s *DeclarationRecordTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitDeclarationRecordType(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclarationArrayTypeContext struct {
	*DeclarationType2Context
}

func NewDeclarationArrayTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclarationArrayTypeContext {
	var p = new(DeclarationArrayTypeContext)

	p.DeclarationType2Context = NewEmptyDeclarationType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationType2Context))

	return p
}

func (s *DeclarationArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationArrayTypeContext) Identifiant() IIdentifiantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifiantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifiantContext)
}

func (s *DeclarationArrayTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitDeclarationArrayType(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclarationTypeClassiqueContext struct {
	*DeclarationType2Context
}

func NewDeclarationTypeClassiqueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclarationTypeClassiqueContext {
	var p = new(DeclarationTypeClassiqueContext)

	p.DeclarationType2Context = NewEmptyDeclarationType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationType2Context))

	return p
}

func (s *DeclarationTypeClassiqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationTypeClassiqueContext) Identifiant() IIdentifiantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifiantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifiantContext)
}

func (s *DeclarationTypeClassiqueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitDeclarationTypeClassique(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) DeclarationType2() (localctx IDeclarationType2Context) {
	this := p
	_ = this

	localctx = NewDeclarationType2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, tigerParserRULE_declarationType2)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(258)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tigerParserID:
		localctx = NewDeclarationTypeClassiqueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(242)
			p.Identifiant()
		}

	case tigerParserT__34:
		localctx = NewDeclarationArrayTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(243)
			p.Match(tigerParserT__34)
		}
		{
			p.SetState(244)
			p.Match(tigerParserT__19)
		}
		{
			p.SetState(245)
			p.Identifiant()
		}

	case tigerParserT__21:
		localctx = NewDeclarationRecordTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(246)
			p.Match(tigerParserT__21)
		}
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == tigerParserID {
			{
				p.SetState(247)
				p.DeclarationChamp()
			}
			p.SetState(252)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == tigerParserT__16 {
				{
					p.SetState(248)
					p.Match(tigerParserT__16)
				}
				{
					p.SetState(249)
					p.DeclarationChamp()
				}

				p.SetState(254)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(257)
			p.Match(tigerParserT__22)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDeclarationChampContext is an interface to support dynamic dispatch.
type IDeclarationChampContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifiant() []IIdentifiantContext
	Identifiant(i int) IIdentifiantContext

	// IsDeclarationChampContext differentiates from other interfaces.
	IsDeclarationChampContext()
}

type DeclarationChampContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationChampContext() *DeclarationChampContext {
	var p = new(DeclarationChampContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_declarationChamp
	return p
}

func (*DeclarationChampContext) IsDeclarationChampContext() {}

func NewDeclarationChampContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationChampContext {
	var p = new(DeclarationChampContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_declarationChamp

	return p
}

func (s *DeclarationChampContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationChampContext) AllIdentifiant() []IIdentifiantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifiantContext); ok {
			len++
		}
	}

	tst := make([]IIdentifiantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifiantContext); ok {
			tst[i] = t.(IIdentifiantContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationChampContext) Identifiant(i int) IIdentifiantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifiantContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifiantContext)
}

func (s *DeclarationChampContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationChampContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationChampContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitDeclarationChamp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) DeclarationChamp() (localctx IDeclarationChampContext) {
	this := p
	_ = this

	localctx = NewDeclarationChampContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, tigerParserRULE_declarationChamp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		p.Identifiant()
	}
	{
		p.SetState(261)
		p.Match(tigerParserT__35)
	}
	{
		p.SetState(262)
		p.Identifiant()
	}

	return localctx
}

// IDeclarationFonctionContext is an interface to support dynamic dispatch.
type IDeclarationFonctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifiant() []IIdentifiantContext
	Identifiant(i int) IIdentifiantContext
	Expression() IExpressionContext
	AllDeclarationChamp() []IDeclarationChampContext
	DeclarationChamp(i int) IDeclarationChampContext

	// IsDeclarationFonctionContext differentiates from other interfaces.
	IsDeclarationFonctionContext()
}

type DeclarationFonctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationFonctionContext() *DeclarationFonctionContext {
	var p = new(DeclarationFonctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_declarationFonction
	return p
}

func (*DeclarationFonctionContext) IsDeclarationFonctionContext() {}

func NewDeclarationFonctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationFonctionContext {
	var p = new(DeclarationFonctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_declarationFonction

	return p
}

func (s *DeclarationFonctionContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationFonctionContext) AllIdentifiant() []IIdentifiantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifiantContext); ok {
			len++
		}
	}

	tst := make([]IIdentifiantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifiantContext); ok {
			tst[i] = t.(IIdentifiantContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationFonctionContext) Identifiant(i int) IIdentifiantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifiantContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifiantContext)
}

func (s *DeclarationFonctionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeclarationFonctionContext) AllDeclarationChamp() []IDeclarationChampContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationChampContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationChampContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationChampContext); ok {
			tst[i] = t.(IDeclarationChampContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationFonctionContext) DeclarationChamp(i int) IDeclarationChampContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationChampContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationChampContext)
}

func (s *DeclarationFonctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationFonctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationFonctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitDeclarationFonction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) DeclarationFonction() (localctx IDeclarationFonctionContext) {
	this := p
	_ = this

	localctx = NewDeclarationFonctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, tigerParserRULE_declarationFonction)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.Match(tigerParserT__36)
	}
	{
		p.SetState(265)
		p.Identifiant()
	}
	{
		p.SetState(266)
		p.Match(tigerParserT__13)
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tigerParserID {
		{
			p.SetState(267)
			p.DeclarationChamp()
		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == tigerParserT__16 {
			{
				p.SetState(268)
				p.Match(tigerParserT__16)
			}
			{
				p.SetState(269)
				p.DeclarationChamp()
			}

			p.SetState(274)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(277)
		p.Match(tigerParserT__15)
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tigerParserT__35 {
		{
			p.SetState(278)
			p.Match(tigerParserT__35)
		}
		{
			p.SetState(279)
			p.Identifiant()
		}

	}
	{
		p.SetState(282)
		p.Match(tigerParserT__3)
	}
	{
		p.SetState(283)
		p.Expression()
	}

	return localctx
}

// IDeclarationValeurContext is an interface to support dynamic dispatch.
type IDeclarationValeurContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifiant() []IIdentifiantContext
	Identifiant(i int) IIdentifiantContext
	Expression() IExpressionContext

	// IsDeclarationValeurContext differentiates from other interfaces.
	IsDeclarationValeurContext()
}

type DeclarationValeurContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationValeurContext() *DeclarationValeurContext {
	var p = new(DeclarationValeurContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_declarationValeur
	return p
}

func (*DeclarationValeurContext) IsDeclarationValeurContext() {}

func NewDeclarationValeurContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationValeurContext {
	var p = new(DeclarationValeurContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_declarationValeur

	return p
}

func (s *DeclarationValeurContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationValeurContext) AllIdentifiant() []IIdentifiantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifiantContext); ok {
			len++
		}
	}

	tst := make([]IIdentifiantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifiantContext); ok {
			tst[i] = t.(IIdentifiantContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationValeurContext) Identifiant(i int) IIdentifiantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifiantContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifiantContext)
}

func (s *DeclarationValeurContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeclarationValeurContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationValeurContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationValeurContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitDeclarationValeur(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) DeclarationValeur() (localctx IDeclarationValeurContext) {
	this := p
	_ = this

	localctx = NewDeclarationValeurContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, tigerParserRULE_declarationValeur)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.Match(tigerParserT__37)
	}
	{
		p.SetState(286)
		p.Identifiant()
	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tigerParserT__35 {
		{
			p.SetState(287)
			p.Match(tigerParserT__35)
		}
		{
			p.SetState(288)
			p.Identifiant()
		}

	}
	{
		p.SetState(291)
		p.Match(tigerParserT__0)
	}
	{
		p.SetState(292)
		p.Expression()
	}

	return localctx
}

// IConstantesContext is an interface to support dynamic dispatch.
type IConstantesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsConstantesContext differentiates from other interfaces.
	IsConstantesContext()
}

type ConstantesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantesContext() *ConstantesContext {
	var p = new(ConstantesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_constantes
	return p
}

func (*ConstantesContext) IsConstantesContext() {}

func NewConstantesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantesContext {
	var p = new(ConstantesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_constantes

	return p
}

func (s *ConstantesContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantesContext) CopyFrom(ctx *ConstantesContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ConstantesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ChaineChrContext struct {
	*ConstantesContext
}

func NewChaineChrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ChaineChrContext {
	var p = new(ChaineChrContext)

	p.ConstantesContext = NewEmptyConstantesContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantesContext))

	return p
}

func (s *ChaineChrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChaineChrContext) STR() antlr.TerminalNode {
	return s.GetToken(tigerParserSTR, 0)
}

func (s *ChaineChrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitChaineChr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilContext struct {
	*ConstantesContext
}

func NewNilContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilContext {
	var p = new(NilContext)

	p.ConstantesContext = NewEmptyConstantesContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantesContext))

	return p
}

func (s *NilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitNil(s)

	default:
		return t.VisitChildren(s)
	}
}

type BreakContext struct {
	*ConstantesContext
}

func NewBreakContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakContext {
	var p = new(BreakContext)

	p.ConstantesContext = NewEmptyConstantesContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantesContext))

	return p
}

func (s *BreakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitBreak(s)

	default:
		return t.VisitChildren(s)
	}
}

type EntierContext struct {
	*ConstantesContext
}

func NewEntierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EntierContext {
	var p = new(EntierContext)

	p.ConstantesContext = NewEmptyConstantesContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantesContext))

	return p
}

func (s *EntierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntierContext) INT() antlr.TerminalNode {
	return s.GetToken(tigerParserINT, 0)
}

func (s *EntierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitEntier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) Constantes() (localctx IConstantesContext) {
	this := p
	_ = this

	localctx = NewConstantesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, tigerParserRULE_constantes)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(298)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tigerParserSTR:
		localctx = NewChaineChrContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(294)
			p.Match(tigerParserSTR)
		}

	case tigerParserINT:
		localctx = NewEntierContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(295)
			p.Match(tigerParserINT)
		}

	case tigerParserT__38:
		localctx = NewNilContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(296)
			p.Match(tigerParserT__38)
		}

	case tigerParserT__39:
		localctx = NewBreakContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(297)
			p.Match(tigerParserT__39)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIdentifiantContext is an interface to support dynamic dispatch.
type IIdentifiantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsIdentifiantContext differentiates from other interfaces.
	IsIdentifiantContext()
}

type IdentifiantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifiantContext() *IdentifiantContext {
	var p = new(IdentifiantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tigerParserRULE_identifiant
	return p
}

func (*IdentifiantContext) IsIdentifiantContext() {}

func NewIdentifiantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifiantContext {
	var p = new(IdentifiantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tigerParserRULE_identifiant

	return p
}

func (s *IdentifiantContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifiantContext) ID() antlr.TerminalNode {
	return s.GetToken(tigerParserID, 0)
}

func (s *IdentifiantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifiantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifiantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tigerVisitor:
		return t.VisitIdentifiant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tigerParser) Identifiant() (localctx IIdentifiantContext) {
	this := p
	_ = this

	localctx = NewIdentifiantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, tigerParserRULE_identifiant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		p.Match(tigerParserID)
	}

	return localctx
}
