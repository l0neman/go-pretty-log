package main

import (
	"fmt"
	"time"

	plog "github.com/l0neman/go-pretty-log"
	"github.com/l0neman/go-pretty-log/tool/highlignt"
	"github.com/l0neman/go-pretty-log/tool/table"
)

const logTag = "log_test"

func TestHighlightLine() {
	fmt.Println("--== TestHighlightLine ==--")
	fmt.Println(highlignt.GetLine("欢迎进入 V1.0 系统", 30))

	lines := []string{"欢迎进入 V1.0 系统", "运行中…"}
	fmt.Println(highlignt.GetLines(lines, 25))
}

// Printer 自定义输出
type Printer struct {
	// 如果只想对日志进行处理，可以保留默认实现
	plog.Printer
}

func (p *Printer) Print(time time.Time, level plog.Level, logTag, logContent string, pid int, colorful bool,
	stackInfo string) {

	levelTag := "Normal"
	if level >= plog.LevelWarn {
		levelTag = "Error"
	}

	// 1732532821260 Error >>log_test>> This is a warning level log. x log_test.go:43
	fmt.Printf("%d %s >>%s>> %s x %s", time.UnixMilli(), levelTag, logTag, logContent, stackInfo)

	// 调用默认的实现
	// p.Printer.Print(time, level, logTag, logContent, pid, colorful, stackInfo)
}

func NewPrinter() *Printer {
	return &Printer{plog.NewPrinter()}
}

func TestLog() {
	fmt.Println("--== TestLog ==--")

	// 设置只输出 warn 和 error 级别的日志
	plog.SetLevel(plog.LevelWarn | plog.LevelError)

	// 使用全局日志对象打印
	plog.I(logTag, "This is an info level log.")
	plog.D(logTag, "This is a debug level log.")
	plog.W(logTag, "This is a warning level log.")
	plog.E(logTag, "This is a error level log.")
	// plog.Fatalln("log_test", "This is a fatal level log.")
	// plog.Panicln(logTag, "This is a panic level log.")

	// 使用局部日志对象打印
	localLogger := plog.NewLogger()
	localLogger.SetFlag(plog.FlagStackEnabled)
	localLogger.I(logTag, "This is a custom info level log.")
}

func TestCustomPrint() {
	// 自定义输出器
	plog.GlobalLogger().SetPrinter(NewPrinter())

	plog.I(logTag, "This is an info level log.")
	plog.D(logTag, "This is a debug level log.")
	plog.W(logTag, "This is a warning level log.")
	plog.E(logTag, "This is a error level log.")
}

func TestPrettyTable() {
	fmt.Println("--== TestPrettyTable ==--")
	// 直接获得表格
	content := [][]interface{}{
		{"Name", "Age", "City", "High"},
		{"Alice", 25, "Beijing", "170cm"},
		{"Bob", 30, "San Francisco", "180cm"},
	}

	fmt.Println(table.GetHorizontalPrettyTable(content))

	// 带名称
	fmt.Println(table.GetHorizontalPrettyTableWithName(content, "Members"))

	// 逐行记录表格，统一获得
	prettyTable := table.NewPrettyTable()
	prettyTable.SetGravity(table.GravityHorizontal)
	prettyTable.SetTableName("Members")
	prettyTable.SetTitles("Name", "Age", "City", "High")
	prettyTable.AddValues("Alice", 25, "Beijing", "170cm")
	prettyTable.AddValues("Bob", 30, "San Francisco", "180cm")
	fmt.Println(prettyTable.Get())

	// 垂直表格
	verticalTable := table.NewPrettyTable()
	verticalTable.SetGravity(table.GravityVertical)
	verticalTable.SetTableName("Members")
	verticalTable.SetTitles("Name", "Age", "City", "High")
	verticalTable.AddValues("Alice", 25, "Beijing", "170cm")
	verticalTable.AddValues("Bob", 30, "San Francisco", "180cm")
	fmt.Println(verticalTable.Get())
}

func main() {
	TestHighlightLine()
	TestLog()
	TestPrettyTable()
}
