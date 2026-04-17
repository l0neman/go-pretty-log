package pretty_log

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	// D/Hello: This is a debug level log.
	logTextFormat = "%s/%s: %s"
	// 2024/07/22 19:11:22 11243 D/Hello: This is a debug level log. > hello.go:33
	logFormat = "%s %d %s%s"

	timeFormat = "2006/01/02 15:04:05"

	labelInfo  = "I"
	labelDebug = "D"
	labelWarn  = "W"
	labelError = "E"
	labelFatal = "F"
	labelPanic = "P"
)

var (
	colorInfo  = []int{3, 169, 244} // blue
	colorDebug = []int{76, 175, 80} // green
	colorWarn  = []int{255, 152, 0} // orange
	colorError = []int{244, 67, 54} // red
	colorFatal = []int{121, 85, 72} // brown
	colorPanic = []int{121, 85, 72} // brown

	goLog = log.New(os.Stdout, "", 0)
)

type levelInfo struct {
	color  []int
	label  string
	handle func(string)
}

type PrinterImpl struct {
	levelInfoMap map[Level]levelInfo
}

func (g *PrinterImpl) Print(time time.Time, level Level, logTag, logText string, pid int, colorful bool, stackInfo string) {
	li, _ := g.levelInfoMap[level]

	text := g.buildLog(time, &li, logTag, logText, pid, colorful, stackInfo)
	li.handle(text)
}

func (g *PrinterImpl) buildLog(time time.Time, li *levelInfo, logTag, logText string, pid int, colorful bool, stackInfo string) string {
	text := fmt.Sprintf(logTextFormat, li.label, logTag, logText)

	if colorful {
		text = getColorfulText(li.color, text)
	}

	if stackInfo != "" {
		stackInfo = " < " + stackInfo
	}

	return fmt.Sprintf(logFormat, time.Format(timeFormat), pid, text, stackInfo)
}

func NewPrinter() Printer {
	levelInfoMap := map[Level]levelInfo{
		LevelInfo:  {color: colorInfo, label: labelInfo, handle: logPrint},
		LevelDebug: {color: colorDebug, label: labelDebug, handle: logPrint},
		LevelWarn:  {color: colorWarn, label: labelWarn, handle: logPrint},
		LevelError: {color: colorError, label: labelError, handle: logPrint},
		LevelFatal: {color: colorFatal, label: labelFatal, handle: logFatal},
		LevelPanic: {color: colorPanic, label: labelPanic, handle: logPanic},
	}

	return &PrinterImpl{
		levelInfoMap: levelInfoMap,
	}
}

func logFatal(text string) {
	goLog.Fatal(text)
}

func logPanic(text string) {
	goLog.Panic(text)
}

func logPrint(text string) {
	_, _ = os.Stdout.WriteString(text)
}

func getColorfulText(color []int, text string) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", color[0], color[1], color[2], text)
}
