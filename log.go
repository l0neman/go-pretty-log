package pretty_log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	baseStackOffset = 4 // 栈偏移
)

type loggerImpl struct {
	flag        Flag
	level       Level
	stackOffset int
	printer     Printer
}

func NewLogger() Logger {
	return &loggerImpl{
		level:   LevelInfo | LevelDebug | LevelWarn | LevelError | LevelFatal | LevelPanic,
		flag:    FlagColorEnabled | FlagStackEnabled,
		printer: NewPrinter(),
	}
}

func (l *loggerImpl) SetStackOffset(stackOffset int) {
	l.stackOffset = stackOffset
}

func (l *loggerImpl) SetFlag(flag Flag) {
	l.flag = flag
}

func (l *loggerImpl) SetLevel(level Level) {
	l.level = level
}

func (l *loggerImpl) AddFlag(flag Flag) {
	l.flag |= flag
}

func (l *loggerImpl) SetPrinter(printer Printer) {
	if printer == nil {
		return
	}

	l.printer = printer
}

func (l *loggerImpl) I(tag string, a ...any) {
	l.println(LevelInfo, tag, strings.TrimSpace(fmt.Sprintln(a...)))
}

func (l *loggerImpl) If(tag string, format string, a ...any) {
	l.println(LevelInfo, tag, format, a...)
}

func (l *loggerImpl) D(tag string, a ...any) {
	l.println(LevelDebug, tag, strings.TrimSpace(fmt.Sprintln(a...)))
}

func (l *loggerImpl) Df(tag string, format string, a ...any) {
	l.println(LevelDebug, tag, format, a...)
}

func (l *loggerImpl) W(tag string, a ...any) {
	l.println(LevelWarn, tag, strings.TrimSpace(fmt.Sprintln(a...)))
}

func (l *loggerImpl) Wf(tag string, format string, a ...any) {
	l.println(LevelWarn, tag, format, a...)
}

func (l *loggerImpl) E(tag string, a ...any) {
	l.println(LevelError, tag, strings.TrimSpace(fmt.Sprintln(a...)))
}

func (l *loggerImpl) Ef(tag string, format string, a ...any) {
	l.println(LevelError, tag, format, a...)
}

func (l *loggerImpl) Fatalln(tag string, a ...any) {
	l.println(LevelFatal, tag, strings.TrimSpace(fmt.Sprintln(a...)))
}

func (l *loggerImpl) Fatalf(tag string, format string, a ...any) {
	l.println(LevelFatal, tag, format, a...)
}

func (l *loggerImpl) Panicln(tag string, a ...any) {
	l.println(LevelPanic, tag, strings.TrimSpace(fmt.Sprintln(a...)))
}

func (l *loggerImpl) Panicf(tag string, format string, a ...any) {
	l.println(LevelPanic, tag, format, a...)
}

func (l *loggerImpl) println(level Level, tag, format string, a ...any) {
	if l.level&level == 0 {
		return
	}

	colorful := l.flag&FlagColorEnabled != 0
	stackInfo := ""
	if l.flag&FlagStackEnabled != 0 {
		stackInfo = getStackInfo(baseStackOffset + l.stackOffset)
	}

	l.printer.Print(time.Now(), level, tag, fmt.Sprintf(format, a...), os.Getpid(), colorful, stackInfo)
}

func getStackInfo(stackOffset int) string {
	var pcs [1]uintptr
	runtime.Callers(stackOffset, pcs[:])
	frames := runtime.CallersFrames([]uintptr{pcs[0]})
	frame, _ := frames.Next()
	return fmt.Sprintf("%s:%d\n", filepath.Base(frame.File), frame.Line)
}
