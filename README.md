# Pretty Log

中文 | [English](./README_en.md)

简单美观的日志库。

## 安装

```shell
go get github.com/l0neman/go-pretty-log
```

```go
import (
    plog "github.com/l0neman/go-pretty-log"
)
```

## 输出不同级别日志

```go
// 根据模块确定日志标签
const logTag = "log_test"

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
```

```shell
2024/07/25 20:03:29 1308849 I/log_test: This is an info level log. < log_test.go:25
2024/07/25 20:03:29 1308849 D/log_test: This is a debug level log. < log_test.go:26
2024/07/25 20:03:29 1308849 W/log_test: This is a warning level log. < log_test.go:27
2024/07/25 20:03:29 1308849 E/log_test: This is a error level log. < log_test.go:28
2024/07/25 20:03:29 1308849 I/log_test: This is a custom info level log. < log_test.go:35
```

- 设置输出级别

```go
// 设置只输出 warn 和 error 级别的日志
plog.SetLevel(plog.LevelWarn | plog.LevelError)
```

- 开启颜色（不能保证所有终端都支持）

```go
plog.SetFlag(plog.FlagColorEnabled)
```

![colorful_log.png](./arts/colorful_log.png)

## 自定义输出器

```go
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

// 自定义输出器
plog.GlobalLogger().SetPrinter(NewPrinter())

plog.I(logTag, "This is an info level log.")
plog.D(logTag, "This is a debug level log.")
plog.W(logTag, "This is a warning level log.")
plog.E(logTag, "This is a error level log.")
```

![image-20241125195547950](./arts/custom_print.png)

## 日志工具

### 输出醒目高亮的信息

```go
fmt.Println("--== TestHighlightLine ==--")
fmt.Println(highlignt.GetLine("欢迎进入 V1.0 系统", 30))

lines := []string{"欢迎进入 V1.0 系统", "运行中…"}
fmt.Println(highlignt.GetLines(lines, 25))
```

```shell
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
┃ 欢迎进入 V1.0 系统
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

┏━━━━━━━━━━━━━━━━━━━━━━━━
┃ 欢迎进入 V1.0 系统
┃ 运行中…
┗━━━━━━━━━━━━━━━━━━━━━━━━
```

### 输出表格

不支持中文，因为无法确保对齐。

```go
// 直接获得表格
content := [][]interface{}{
    {"Name", "Age", "City", "High"},
    {"Alice", 25, "Beijing", "170cm"},
    {"Bob", 30, "San Francisco", "180cm"},
}

fmt.Println(plog.GetHorizontalPrettyTable(content))

// 带名称
fmt.Println(plog.GetHorizontalPrettyTableWithName(content, "Members"))
```

```shell
┌──────────────────────────────────┐
│ Name   Age  City           High  │
│ ─────  ───  ─────────────  ───── │
│ Alice  25   Beijing        170cm │
│ Bob    30   San Francisco  180cm │
└──────────────────────────────────┘
┌──────────────────────────────────┐
│ Members                          │
├──────────────────────────────────┤
│ Name   Age  City           High  │
│ ─────  ───  ─────────────  ───── │
│ Alice  25   Beijing        170cm │
│ Bob    30   San Francisco  180cm │
└──────────────────────────────────┘
```

以创建对象的方式输出水平表格。

```go
// 逐行记录表格，统一获得
prettyTable := plog.NewPrettyTable()
prettyTable.SetGravity(plog.GravityHorizontal)
prettyTable.SetTableName("Members")
prettyTable.SetTitles("Name", "Age", "City", "High")
prettyTable.AddValues("Alice", 25, "Beijing", "170cm")
prettyTable.AddValues("Bob", 30, "San Francisco", "180cm")
fmt.Println(prettyTable.Get())
```

```shell
┌──────────────────────────────────┐
│ Members                          │
├──────────────────────────────────┤
│ Name   Age  City           High  │
│ ─────  ───  ─────────────  ───── │
│ Alice  25   Beijing        170cm │
│ Bob    30   San Francisco  180cm │
└──────────────────────────────────┘
```

如果表格列过多，导致折行 ，可选择输出垂直表格。

```go
// 垂直表格
verticalTable := plog.NewPrettyTable()
verticalTable.SetGravity(plog.GravityVertical)
verticalTable.SetTableName("Members")
verticalTable.SetTitles("Name", "Age", "City", "High")
verticalTable.AddValues("Alice", 25, "Beijing", "170cm")
verticalTable.AddValues("Bob", 30, "San Francisco", "180cm")
fmt.Println(verticalTable.Get())
```

```shell
┌────────────────────╼
│       Members       
├────────[ 0 ]───────┈
│ Name: Alice
│  Age: 25
│ City: Beijing
│ High: 170cm
├────────[ 1 ]───────┈
│ Name: Bob
│  Age: 30
│ City: San Francisco
│ High: 180cm
└────────────────────╼
```