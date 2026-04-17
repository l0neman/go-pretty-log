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
	fmt.Println(highlignt.GetLine("Welcome to V1.0 System", 30))

	lines := []string{"Welcome to V1.0 System", "Running..."}
	fmt.Println(highlignt.GetLines(lines, 25))
}

// Printer custom printer
type Printer struct {
	// If you only want to process logs, you can keep the default implementation.
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

	// Call the default implementation
	// p.Printer.Print(time, level, logTag, logContent, pid, colorful, stackInfo)
}

func NewPrinter() *Printer {
	return &Printer{plog.NewPrinter()}
}

func TestLog() {
	fmt.Println("--== TestLog ==--")

	// Set to only output logs of warn and error levels
	plog.SetLevel(plog.LevelWarn | plog.LevelError)

	// 使用全局日志对象打印
	plog.I(logTag, "This is an info level log.")
	plog.D(logTag, "This is a debug level log.")
	plog.W(logTag, "This is a warning level log.")
	plog.E(logTag, "This is a error level log.")
	// plog.Fatalln("log_test", "This is a fatal level log.")
	// plog.Panicln(logTag, "This is a panic level log.")

	// Print using a local log object
	localLogger := plog.NewLogger()
	localLogger.SetFlag(plog.FlagStackEnabled)
	localLogger.I(logTag, "This is a custom info level log.")
}

func TestCustomPrint() {
	// Custom printer
	plog.GlobalLogger().SetPrinter(NewPrinter())

	plog.I(logTag, "This is an info level log.")
	plog.D(logTag, "This is a debug level log.")
	plog.W(logTag, "This is a warning level log.")
	plog.E(logTag, "This is a error level log.")
}

func TestPrettyTable() {
	fmt.Println("--== TestPrettyTable ==--")
	// Get table directly
	content := [][]interface{}{
		{"Name", "Age", "City", "High"},
		{"Alice", 25, "Beijing", "170cm"},
		{"Bob", 30, "San Francisco", "180cm"},
	}

	fmt.Println(table.GetHorizontalPrettyTable(content))

	// With name
	fmt.Println(table.GetHorizontalPrettyTableWithName(content, "Members"))

	// Record the table row by row, and obtain it in a unified manner
	prettyTable := table.NewPrettyTable()
	prettyTable.SetGravity(table.GravityHorizontal)
	prettyTable.SetTableName("Members")
	prettyTable.SetTitles("Name", "Age", "City", "High")
	prettyTable.AddValues("Alice", 25, "Beijing", "170cm")
	prettyTable.AddValues("Bob", 30, "San Francisco", "180cm")
	fmt.Println(prettyTable.Get())

	// Vertical table
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
