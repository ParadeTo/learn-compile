package playscript

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

const (
	CompilationLogINFO = iota
	CompilationLogWARNING
	CompilationLogERROR
)

// 记录编译过程中产生的信息
type CompilationLog struct {
	message        string
	line           int
	positionInLine int
	ctx            antlr.ParserRuleContext
	_type          int
}

func (l *CompilationLog) ToString() string {
	return fmt.Sprintf("%s @%d:%d", l.message, l.line, l.positionInLine)
}

func NewCompilationLog() *CompilationLog {
	return &CompilationLog{_type: CompilationLogINFO}
}
