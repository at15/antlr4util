package antlr4util

import (
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// CaseInsensitiveStream converts all token to UPPERCASE for lexer.
// It allows us to write lexer rules in UPPERCASE but accepts mixedCASE input and preserve the case.
// It is based on https://github.com/antlr/antlr4/blob/master/doc/resources/case_changing_stream.go.
type CaseInsensitiveStream struct {
	antlr.CharStream
}

// NewCaseInsensitiveStream wraps an underlying input stream to emit UPPERCASE token for lexer.
func NewCaseInsensitiveStream(stream antlr.CharStream) *CaseInsensitiveStream {
	return &CaseInsensitiveStream{stream}
}

// LA converts all token to upper case.
func (ci *CaseInsensitiveStream) LA(offset int) int {
	c := ci.CharStream.LA(offset)
	return int(unicode.ToUpper(rune(c)))
}
