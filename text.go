package antlr4util

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Text returns string with each token separated by space.
// It is based on antlr.BaseParserRuleContext.GetText, which concats children directly w/o sep.
func Text(r antlr.Tree) string {
	// terminal node does not have child
	term, ok := r.(antlr.TerminalNode)
	if ok {
		return term.GetText()
	}

	if r.GetChildCount() == 0 {
		return ""
	}

	var sb strings.Builder
	for _, child := range r.GetChildren() {
		sb.WriteString(Text(child))
		// TODO: for symbol like , {, we need to trim extra space
		sb.WriteRune(' ')
	}
	return sb.String()
}
