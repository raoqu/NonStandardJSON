package main

import (
	"fmt"
	"strings"

	"NJson/parser"

	"github.com/antlr4-go/antlr/v4"
)

const DEBUG_EVERY_RULE = false

type NJsonFormatListener struct {
	indentLevel int
}

func NewNJsonFormatListener() *NJsonFormatListener {
	return &NJsonFormatListener{
		indentLevel: 0,
	}
}

// VisitTerminal is called when a terminal node is visited.
func (s *NJsonFormatListener) VisitTerminal(node antlr.TerminalNode) {
	s.output("VisitTerminal - "+node.GetText(), 0)
}

// VisitErrorNode is called when an error node is visited.
func (s *NJsonFormatListener) VisitErrorNode(node antlr.ErrorNode) {
	token := node.GetSymbol()
	line := token.GetLine()
	column := token.GetColumn()
	s.output(fmt.Sprintf("VisitErrorNode - ERROR %d,%d", line, column), 0)
}

// EnterEveryRule is called when any rule is entered.
func (s *NJsonFormatListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	if DEBUG_EVERY_RULE {
		s.output("EnterEveryRule", 1)
		// indent := strings.Repeat("  ", s.indentLevel)

		switch ctx.GetRuleContext().(type) {
		case *parser.PairContext:
			s.output("pair", 0)
			// pair := ctx.(*parser.PairContext)
			// fmt.Printf("%sKey: %s\n", indent, pair.STRING().GetText())
			// fmt.Printf("%sValue: %s\n", indent, pair.Value().GetText())
		case *parser.ArrayContext:
			s.output("value", 0)
			// array := ctx.(*parser.ArrayContext)
			// fmt.Printf("%sArray Start\n", indent)
			// for _, child := range array.GetChildren() {
			// 	child.GetChild(0)
			// 	fmt.Printf("%sArray Item: %s\n", indent, "") //child.GetText())
			// }
			// fmt.Printf("%sArray End\n", indent)
		case *parser.ObjectContext:
			s.output("object", 0)
			// obj := ctx.(*parser.ObjectContext)
			// for i, child := range obj.GetChildren() {
			// 	child.GetChild(0)
			// 	fmt.Printf("Item: %s\n", i) //child.GetText())
			// }
		case *parser.ValueContext:
			s.output("value", 0)
		}
	}
}

// ExitEveryRule is called when any rule is exited.
func (s *NJsonFormatListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	if DEBUG_EVERY_RULE {
		s.output("ExitEveryRule", -1)
	}
}

// EnterJson is called when production json is entered.
func (s *NJsonFormatListener) EnterJson(ctx *parser.JsonContext) {
	s.output("EnterJson", 1)
}

// ExitJson is called when production json is exited.
func (s *NJsonFormatListener) ExitJson(ctx *parser.JsonContext) {
	s.output("ExitJson", -1)
}

// EnterObject is called when production object is entered.
func (s *NJsonFormatListener) EnterObject(ctx *parser.ObjectContext) {
	s.output("EnterObject", 1)
}

// ExitObject is called when production object is exited.
func (s *NJsonFormatListener) ExitObject(ctx *parser.ObjectContext) {
	s.output("ExitObject", -1)
}

// EnterPair is called when production pair is entered.
func (s *NJsonFormatListener) EnterPair(ctx *parser.PairContext) {
	s.output("EnterPair", 1)
}

// ExitPair is called when production pair is exited.
func (s *NJsonFormatListener) ExitPair(ctx *parser.PairContext) {
	s.output("ExitPair", -1)
}

// EnterArray is called when production array is entered.
func (s *NJsonFormatListener) EnterArray(ctx *parser.ArrayContext) {
	s.output("EnterArray", 1)
}

// ExitArray is called when production array is exited.
func (s *NJsonFormatListener) ExitArray(ctx *parser.ArrayContext) {
	s.output("ExitArray", -1)
}

// EnterValue is called when production value is entered.
func (s *NJsonFormatListener) EnterValue(ctx *parser.ValueContext) {
	val := ""
	switch {
	case ctx.STRING() != nil:
		val = fmt.Sprintf("STRING, %s", ctx.STRING().GetSymbol().GetText())
	case ctx.NUMBER() != nil:
		val = fmt.Sprintf("NUMBER, %s", ctx.NUMBER().GetText())
	case ctx.GetText() == "true":
		val = "BOOLEAN, true"
	case ctx.GetText() == "false":
		val = "BOOLEAN, false"
	case ctx.GetText() == "null":
		val = "NULL, null"
	case ctx.Object() != nil:
		val = fmt.Sprintf("OBJECT, %s", ctx.Object().GetText())
	case ctx.Array() != nil:
		val = fmt.Sprintf("ARRAY, %s", ctx.Array().GetText())
	}
	s.output("Value - "+val, 1)
}

// ExitValue is called when production value is exited.
func (s *NJsonFormatListener) ExitValue(ctx *parser.ValueContext) {
	s.output("", -1)
}

func (f *NJsonFormatListener) printIndentation() {
	fmt.Println(strings.Repeat("..", f.indentLevel))
}

func (f *NJsonFormatListener) output(s string, i int) {
	prefix := "* "
	indent := 0
	if i == 0 {
		prefix = ""
		indent = 1
	}
	if i > 0 {
		f.indentLevel += i
	}
	if s != "" {
		print(strings.Repeat("  ", f.indentLevel+indent) + prefix)
	}
	if i < 0 {
		f.indentLevel += i
	}
	if s != "" {
		println(s)
	}
}
