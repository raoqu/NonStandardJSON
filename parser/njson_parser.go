// Code generated from NJson.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // NJson

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type NJsonParser struct {
	*antlr.BaseParser
}

var NJsonParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func njsonParserInit() {
	staticData := &NJsonParserStaticData
	staticData.LiteralNames = []string{
		"", "'{'", "','", "'}'", "':'", "'['", "']'", "'true'", "'false'", "'null'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "STRING", "NUMBER", "WS",
	}
	staticData.RuleNames = []string{
		"json", "object", "pair", "array", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 12, 58, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 1, 0, 3, 0, 13, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 19, 8, 1,
		10, 1, 12, 1, 22, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 28, 8, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 38, 8, 3, 10, 3, 12, 3, 41,
		9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 47, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 3, 4, 56, 8, 4, 1, 4, 0, 0, 5, 0, 2, 4, 6, 8, 0, 0, 63,
		0, 12, 1, 0, 0, 0, 2, 27, 1, 0, 0, 0, 4, 29, 1, 0, 0, 0, 6, 46, 1, 0, 0,
		0, 8, 55, 1, 0, 0, 0, 10, 13, 3, 2, 1, 0, 11, 13, 3, 6, 3, 0, 12, 10, 1,
		0, 0, 0, 12, 11, 1, 0, 0, 0, 13, 1, 1, 0, 0, 0, 14, 15, 5, 1, 0, 0, 15,
		20, 3, 4, 2, 0, 16, 17, 5, 2, 0, 0, 17, 19, 3, 4, 2, 0, 18, 16, 1, 0, 0,
		0, 19, 22, 1, 0, 0, 0, 20, 18, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 23,
		1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 23, 24, 5, 3, 0, 0, 24, 28, 1, 0, 0, 0,
		25, 26, 5, 1, 0, 0, 26, 28, 5, 3, 0, 0, 27, 14, 1, 0, 0, 0, 27, 25, 1,
		0, 0, 0, 28, 3, 1, 0, 0, 0, 29, 30, 5, 10, 0, 0, 30, 31, 5, 4, 0, 0, 31,
		32, 3, 8, 4, 0, 32, 5, 1, 0, 0, 0, 33, 34, 5, 5, 0, 0, 34, 39, 3, 8, 4,
		0, 35, 36, 5, 2, 0, 0, 36, 38, 3, 8, 4, 0, 37, 35, 1, 0, 0, 0, 38, 41,
		1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 42, 1, 0, 0, 0,
		41, 39, 1, 0, 0, 0, 42, 43, 5, 6, 0, 0, 43, 47, 1, 0, 0, 0, 44, 45, 5,
		5, 0, 0, 45, 47, 5, 6, 0, 0, 46, 33, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 47,
		7, 1, 0, 0, 0, 48, 56, 5, 10, 0, 0, 49, 56, 5, 11, 0, 0, 50, 56, 3, 2,
		1, 0, 51, 56, 3, 6, 3, 0, 52, 56, 5, 7, 0, 0, 53, 56, 5, 8, 0, 0, 54, 56,
		5, 9, 0, 0, 55, 48, 1, 0, 0, 0, 55, 49, 1, 0, 0, 0, 55, 50, 1, 0, 0, 0,
		55, 51, 1, 0, 0, 0, 55, 52, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 54, 1,
		0, 0, 0, 56, 9, 1, 0, 0, 0, 6, 12, 20, 27, 39, 46, 55,
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

// NJsonParserInit initializes any static state used to implement NJsonParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNJsonParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NJsonParserInit() {
	staticData := &NJsonParserStaticData
	staticData.once.Do(njsonParserInit)
}

// NewNJsonParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNJsonParser(input antlr.TokenStream) *NJsonParser {
	NJsonParserInit()
	this := new(NJsonParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &NJsonParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "NJson.g4"

	return this
}

// NJsonParser tokens.
const (
	NJsonParserEOF    = antlr.TokenEOF
	NJsonParserT__0   = 1
	NJsonParserT__1   = 2
	NJsonParserT__2   = 3
	NJsonParserT__3   = 4
	NJsonParserT__4   = 5
	NJsonParserT__5   = 6
	NJsonParserT__6   = 7
	NJsonParserT__7   = 8
	NJsonParserT__8   = 9
	NJsonParserSTRING = 10
	NJsonParserNUMBER = 11
	NJsonParserWS     = 12
)

// NJsonParser rules.
const (
	NJsonParserRULE_json   = 0
	NJsonParserRULE_object = 1
	NJsonParserRULE_pair   = 2
	NJsonParserRULE_array  = 3
	NJsonParserRULE_value  = 4
)

// IJsonContext is an interface to support dynamic dispatch.
type IJsonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Object() IObjectContext
	Array() IArrayContext

	// IsJsonContext differentiates from other interfaces.
	IsJsonContext()
}

type JsonContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonContext() *JsonContext {
	var p = new(JsonContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NJsonParserRULE_json
	return p
}

func InitEmptyJsonContext(p *JsonContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NJsonParserRULE_json
}

func (*JsonContext) IsJsonContext() {}

func NewJsonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonContext {
	var p = new(JsonContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NJsonParserRULE_json

	return p
}

func (s *JsonContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonContext) Object() IObjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *JsonContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *JsonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJsonListener); ok {
		listenerT.EnterJson(s)
	}
}

func (s *JsonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJsonListener); ok {
		listenerT.ExitJson(s)
	}
}

func (p *NJsonParser) Json() (localctx IJsonContext) {
	localctx = NewJsonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NJsonParserRULE_json)
	p.SetState(12)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NJsonParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(10)
			p.Object()
		}

	case NJsonParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(11)
			p.Array()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObjectContext is an interface to support dynamic dispatch.
type IObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPair() []IPairContext
	Pair(i int) IPairContext

	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectContext() *ObjectContext {
	var p = new(ObjectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NJsonParserRULE_object
	return p
}

func InitEmptyObjectContext(p *ObjectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NJsonParserRULE_object
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NJsonParserRULE_object

	return p
}

func (s *ObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectContext) AllPair() []IPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPairContext); ok {
			len++
		}
	}

	tst := make([]IPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPairContext); ok {
			tst[i] = t.(IPairContext)
			i++
		}
	}

	return tst
}

func (s *ObjectContext) Pair(i int) IPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairContext); ok {
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

	return t.(IPairContext)
}

func (s *ObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJsonListener); ok {
		listenerT.EnterObject(s)
	}
}

func (s *ObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJsonListener); ok {
		listenerT.ExitObject(s)
	}
}

func (p *NJsonParser) Object() (localctx IObjectContext) {
	localctx = NewObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NJsonParserRULE_object)
	var _la int

	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(14)
			p.Match(NJsonParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(15)
			p.Pair()
		}
		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == NJsonParserT__1 {
			{
				p.SetState(16)
				p.Match(NJsonParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(17)
				p.Pair()
			}

			p.SetState(22)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(23)
			p.Match(NJsonParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.Match(NJsonParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(26)
			p.Match(NJsonParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	Value() IValueContext

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NJsonParserRULE_pair
	return p
}

func InitEmptyPairContext(p *PairContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NJsonParserRULE_pair
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NJsonParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) STRING() antlr.TerminalNode {
	return s.GetToken(NJsonParserSTRING, 0)
}

func (s *PairContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJsonListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJsonListener); ok {
		listenerT.ExitPair(s)
	}
}

func (p *NJsonParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NJsonParserRULE_pair)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)
		p.Match(NJsonParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(30)
		p.Match(NJsonParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(31)
		p.Value()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllValue() []IValueContext
	Value(i int) IValueContext

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NJsonParserRULE_array
	return p
}

func InitEmptyArrayContext(p *ArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NJsonParserRULE_array
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NJsonParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ArrayContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
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

	return t.(IValueContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJsonListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJsonListener); ok {
		listenerT.ExitArray(s)
	}
}

func (p *NJsonParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NJsonParserRULE_array)
	var _la int

	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Match(NJsonParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.Value()
		}
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == NJsonParserT__1 {
			{
				p.SetState(35)
				p.Match(NJsonParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(36)
				p.Value()
			}

			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(42)
			p.Match(NJsonParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.Match(NJsonParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.Match(NJsonParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	Object() IObjectContext
	Array() IArrayContext

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NJsonParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NJsonParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NJsonParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(NJsonParserSTRING, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(NJsonParserNUMBER, 0)
}

func (s *ValueContext) Object() IObjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *ValueContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJsonListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJsonListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *NJsonParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NJsonParserRULE_value)
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NJsonParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.Match(NJsonParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case NJsonParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.Match(NJsonParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case NJsonParserT__0:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.Object()
		}

	case NJsonParserT__4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(51)
			p.Array()
		}

	case NJsonParserT__6:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(52)
			p.Match(NJsonParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case NJsonParserT__7:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(53)
			p.Match(NJsonParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case NJsonParserT__8:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(54)
			p.Match(NJsonParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
