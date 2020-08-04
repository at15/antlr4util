package antlr4util

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dyweb/gommon/errors"
)

// TODO: not enough error information is extracted, should print actual error location

var _ antlr.ErrorListener = (*ErrorListener)(nil)

// TODO(gommon): generate using stringer
type ErrorType int

const (
	ErrorTypeUnknown = iota
	ErrorTypeSyntax
	ErrorTypeAmbiguity
	ErrorTypeAttemptingFullContext
	ErrorTypeContextSensitivity
)

func (t ErrorType) String() string {
	return []string{
		"unknown",
		"syntax",
		"ambiguity",
		"attempting full context",
		"context sensitivity",
	}[t]
}

type Error struct {
	Type   ErrorType
	Msg    string
	Line   int
	Column int
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s at %d:%d %s", e.Type, e.Line, e.Column, e.Msg)
}

type ErrorListener struct {
	errors errors.MultiErr
}

func NewErrorListener() *ErrorListener {
	return &ErrorListener{
		errors: errors.NewMultiErr(),
	}
}

// e2 == engineering 2
func (e2 ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	err := &Error{
		Type:   ErrorTypeSyntax,
		Line:   line,
		Column: column,
		Msg:    msg,
	}
	e2.errors.Append(err)
}

func (e2 ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	err := &Error{
		Type: ErrorTypeAmbiguity,
		Msg:  recognizer.GetTokenStream().GetTextFromInterval(antlr.NewInterval(startIndex, stopIndex)),
	}
	e2.errors.Append(err)
}

func (e2 ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	err := &Error{
		Type: ErrorTypeAttemptingFullContext,
		Msg:  recognizer.GetTokenStream().GetTextFromInterval(antlr.NewInterval(startIndex, stopIndex)),
	}
	e2.errors.Append(err)
}

func (e2 ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	err := &Error{
		Type: ErrorTypeContextSensitivity,
		Msg:  recognizer.GetTokenStream().GetTextFromInterval(antlr.NewInterval(startIndex, stopIndex)),
	}
	e2.errors.Append(err)
}
