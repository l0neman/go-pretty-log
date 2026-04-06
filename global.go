package pretty_log

var globalLogger Logger

func init() {
	globalLogger = NewLogger()
	globalLogger.SetStackOffset(1)
}

func GlobalLogger() Logger {
	return globalLogger
}

func SetLevel(level Level) {
	globalLogger.SetLevel(level)
}

func I(tag string, a ...any) {
	globalLogger.I(tag, a...)
}

func If(tag string, format string, a ...any) {
	globalLogger.If(tag, format, a...)
}

func D(tag string, a ...any) {
	globalLogger.D(tag, a...)
}

func Df(tag string, format string, a ...any) {
	globalLogger.Df(tag, format, a...)
}

func W(tag string, a ...any) {
	globalLogger.W(tag, a...)
}

func Wf(tag string, format string, a ...any) {
	globalLogger.Wf(tag, format, a...)
}

func E(tag string, a ...any) {
	globalLogger.E(tag, a...)
}

func Ef(tag string, format string, a ...any) {
	globalLogger.Ef(tag, format, a...)
}

func Fatalln(tag string, a ...any) {
	globalLogger.Fatalln(tag, a...)
}

func Fatalf(tag string, format string, a ...any) {
	globalLogger.Fatalf(tag, format, a...)
}

func Panicln(tag string, a ...any) {
	globalLogger.Panicln(tag, a...)
}

func Panicf(tag string, format string, a ...any) {
	globalLogger.Panicf(tag, format, a...)
}
