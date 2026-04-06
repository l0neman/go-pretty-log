package pretty_log

import (
	"time"
)

type Printer interface {
	Print(time time.Time, level Level, logTag, logText string, pid int, colorful bool, stackInfo string)
}

type Logger interface {
	SetStackOffset(int)
	SetFlag(Flag)
	SetLevel(Level)
	SetPrinter(Printer)
	I(tag string, a ...any)
	If(tag string, format string, a ...any)
	D(tag string, a ...any)
	Df(tag string, format string, a ...any)
	W(tag string, a ...any)
	Wf(tag string, format string, a ...any)
	E(tag string, a ...any)
	Ef(tag string, format string, a ...any)
	Fatalln(tag string, a ...any)
	Fatalf(tag string, format string, a ...any)
	Panicln(tag string, a ...any)
	Panicf(tag string, format string, a ...any)
}
