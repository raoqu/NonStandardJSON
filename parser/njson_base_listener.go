// Code generated from NJson.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // NJson

import "github.com/antlr4-go/antlr/v4"

// BaseNJsonListener is a complete listener for a parse tree produced by NJsonParser.
type BaseNJsonListener struct{}

var _ NJsonListener = &BaseNJsonListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNJsonListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNJsonListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNJsonListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNJsonListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterJson is called when production json is entered.
func (s *BaseNJsonListener) EnterJson(ctx *JsonContext) {}

// ExitJson is called when production json is exited.
func (s *BaseNJsonListener) ExitJson(ctx *JsonContext) {}

// EnterObject is called when production object is entered.
func (s *BaseNJsonListener) EnterObject(ctx *ObjectContext) {}

// ExitObject is called when production object is exited.
func (s *BaseNJsonListener) ExitObject(ctx *ObjectContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseNJsonListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseNJsonListener) ExitPair(ctx *PairContext) {}

// EnterArray is called when production array is entered.
func (s *BaseNJsonListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseNJsonListener) ExitArray(ctx *ArrayContext) {}

// EnterValue is called when production value is entered.
func (s *BaseNJsonListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseNJsonListener) ExitValue(ctx *ValueContext) {}
