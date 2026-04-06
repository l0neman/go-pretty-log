package pretty_log

import (
	"io"
	"os"
	"testing"
	"time"
)

// mockPrinter 用于基准测试的 mock 输出器，丢弃所有输出
type mockPrinter struct{}

func (m *mockPrinter) Print(time time.Time, level Level, logTag, logText string, pid int, colorful bool, stackInfo string) {
	// 不做任何输出，模拟最快的打印器
}

// discardPrinter 使用 io.Discard 的打印器
type discardPrinter struct {
	PrinterImpl
}

func newDiscardPrinter() *discardPrinter {
	return &discardPrinter{
		PrinterImpl: PrinterImpl{
			levelInfoMap: map[Level]levelInfo{
				LevelInfo:  {color: colorInfo, label: labelInfo, handle: func(s string) { io.Discard.Write([]byte(s)) }},
				LevelDebug: {color: colorDebug, label: labelDebug, handle: func(s string) { io.Discard.Write([]byte(s)) }},
				LevelWarn:  {color: colorWarn, label: labelWarn, handle: func(s string) { io.Discard.Write([]byte(s)) }},
				LevelError: {color: colorError, label: labelError, handle: func(s string) { io.Discard.Write([]byte(s)) }},
				LevelFatal: {color: colorFatal, label: labelFatal, handle: func(s string) { io.Discard.Write([]byte(s)) }},
				LevelPanic: {color: colorPanic, label: labelPanic, handle: func(s string) { io.Discard.Write([]byte(s)) }},
			},
		},
	}
}

const (
	benchTag     = "benchmark"
	benchMessage = "This is a benchmark test message"
)

// setupLogger 创建一个用于基准测试的 logger
func setupLogger() Logger {
	logger := NewLogger()
	logger.SetPrinter(newDiscardPrinter())
	return logger
}

// BenchmarkLoggerInfo 测试 Info 级别日志性能
func BenchmarkLoggerInfo(b *testing.B) {
	logger := setupLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.I(benchTag, benchMessage)
	}
}

// BenchmarkLoggerInfof 测试格式化 Info 日志性能
func BenchmarkLoggerInfof(b *testing.B) {
	logger := setupLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.If(benchTag, "This is a %s test message %d", "benchmark", i)
	}
}

// BenchmarkLoggerDebug 测试 Debug 级别日志性能
func BenchmarkLoggerDebug(b *testing.B) {
	logger := setupLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.D(benchTag, benchMessage)
	}
}

// BenchmarkLoggerWarn 测试 Warn 级别日志性能
func BenchmarkLoggerWarn(b *testing.B) {
	logger := setupLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.W(benchTag, benchMessage)
	}
}

// BenchmarkLoggerError 测试 Error 级别日志性能
func BenchmarkLoggerError(b *testing.B) {
	logger := setupLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.E(benchTag, benchMessage)
	}
}

// BenchmarkGlobalLoggerInfo 测试全局 logger 的性能
func BenchmarkGlobalLoggerInfo(b *testing.B) {
	oldStdout := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() { os.Stdout = oldStdout }()

	GlobalLogger().SetPrinter(newDiscardPrinter())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		I(benchTag, benchMessage)
	}
}

// BenchmarkLoggerWithColorEnabled 测试启用颜色时的性能
func BenchmarkLoggerWithColorEnabled(b *testing.B) {
	logger := setupLogger()
	logger.SetFlag(FlagColorEnabled)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.I(benchTag, benchMessage)
	}
}

// BenchmarkLoggerWithColorDisabled 测试禁用颜色时的性能
func BenchmarkLoggerWithColorDisabled(b *testing.B) {
	logger := setupLogger()
	logger.SetFlag(FlagClear)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.I(benchTag, benchMessage)
	}
}

// BenchmarkLoggerWithStackEnabled 测试启用栈信息时的性能
func BenchmarkLoggerWithStackEnabled(b *testing.B) {
	logger := setupLogger()
	logger.SetFlag(FlagStackEnabled)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.I(benchTag, benchMessage)
	}
}

// BenchmarkLoggerWithStackDisabled 测试禁用栈信息时的性能
func BenchmarkLoggerWithStackDisabled(b *testing.B) {
	logger := setupLogger()
	logger.SetFlag(FlagClear)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.I(benchTag, benchMessage)
	}
}

// BenchmarkLoggerWithAllFlagsEnabled 测试启用所有 flag 时的性能
func BenchmarkLoggerWithAllFlagsEnabled(b *testing.B) {
	logger := setupLogger()
	logger.SetFlag(FlagColorEnabled | FlagStackEnabled)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.I(benchTag, benchMessage)
	}
}

// BenchmarkLoggerFiltered 测试日志被过滤时的性能
func BenchmarkLoggerFiltered(b *testing.B) {
	logger := setupLogger()
	logger.SetLevel(LevelError) // 只输出 Error 级别，Info 会被过滤
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.I(benchTag, benchMessage) // 这条日志会被过滤掉
	}
}

// BenchmarkLoggerNotFiltered 测试日志未被过滤时的性能
func BenchmarkLoggerNotFiltered(b *testing.B) {
	logger := setupLogger()
	logger.SetLevel(LevelInfo) // 允许 Info 级别
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.I(benchTag, benchMessage)
	}
}

// BenchmarkLoggerMultipleArgs 测试多参数日志的性能
func BenchmarkLoggerMultipleArgs(b *testing.B) {
	logger := setupLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.I(benchTag, "arg1", "arg2", "arg3", 123, 456.789)
	}
}

// BenchmarkLoggerComplexFormatting 测试复杂格式化的性能
func BenchmarkLoggerComplexFormatting(b *testing.B) {
	logger := setupLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.If(benchTag, "User %s logged in at %d with score %.2f from %s", "john_doe", i, 95.5, "192.168.1.1")
	}
}

// BenchmarkLoggerParallel 测试并发日志写入性能
func BenchmarkLoggerParallel(b *testing.B) {
	logger := setupLogger()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.I(benchTag, benchMessage)
		}
	})
}

// BenchmarkLoggerParallelWithAllFlags 测试启用所有 flag 的并发性能
func BenchmarkLoggerParallelWithAllFlags(b *testing.B) {
	logger := setupLogger()
	logger.SetFlag(FlagColorEnabled | FlagStackEnabled)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.I(benchTag, benchMessage)
		}
	})
}

// BenchmarkDifferentLevels 对比不同日志级别的性能
func BenchmarkDifferentLevels(b *testing.B) {
	tests := []struct {
		name string
		fn   func(Logger)
	}{
		{"Info", func(l Logger) { l.I(benchTag, benchMessage) }},
		{"Debug", func(l Logger) { l.D(benchTag, benchMessage) }},
		{"Warn", func(l Logger) { l.W(benchTag, benchMessage) }},
		{"Error", func(l Logger) { l.E(benchTag, benchMessage) }},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			logger := setupLogger()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tt.fn(logger)
			}
		})
	}
}

// BenchmarkFormattingComparison 对比格式化和非格式化性能
func BenchmarkFormattingComparison(b *testing.B) {
	b.Run("NonFormatted", func(b *testing.B) {
		logger := setupLogger()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.I(benchTag, benchMessage)
		}
	})

	b.Run("Formatted", func(b *testing.B) {
		logger := setupLogger()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.If(benchTag, "%s", benchMessage)
		}
	})

	b.Run("FormattedComplex", func(b *testing.B) {
		logger := setupLogger()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.If(benchTag, "message: %s, number: %d", benchMessage, i)
		}
	})
}

// BenchmarkFlagCombinations 测试不同 flag 组合的性能
func BenchmarkFlagCombinations(b *testing.B) {
	tests := []struct {
		name string
		flag Flag
	}{
		{"NoFlags", FlagClear},
		{"ColorOnly", FlagColorEnabled},
		{"StackOnly", FlagStackEnabled},
		{"ColorAndStack", FlagColorEnabled | FlagStackEnabled},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			logger := setupLogger()
			logger.SetFlag(tt.flag)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				logger.I(benchTag, benchMessage)
			}
		})
	}
}

// BenchmarkPrinterImpl 测试默认 Printer 的性能
func BenchmarkPrinterImpl(b *testing.B) {
	printer := newDiscardPrinter()
	now := time.Now()
	pid := os.Getpid()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		printer.Print(now, LevelInfo, benchTag, benchMessage, pid, false, "")
	}
}

// BenchmarkPrinterImplWithColor 测试带颜色的 Printer 性能
func BenchmarkPrinterImplWithColor(b *testing.B) {
	printer := newDiscardPrinter()
	now := time.Now()
	pid := os.Getpid()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		printer.Print(now, LevelInfo, benchTag, benchMessage, pid, true, "")
	}
}

// BenchmarkPrinterImplWithStack 测试带栈信息的 Printer 性能
func BenchmarkPrinterImplWithStack(b *testing.B) {
	printer := newDiscardPrinter()
	now := time.Now()
	pid := os.Getpid()
	stackInfo := "log_test.go:42\n"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		printer.Print(now, LevelInfo, benchTag, benchMessage, pid, false, stackInfo)
	}
}

// BenchmarkGetStackInfo 测试获取栈信息的性能
func BenchmarkGetStackInfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = getStackInfo(baseStackOffset)
	}
}

// BenchmarkNewLogger 测试创建新 Logger 的性能
func BenchmarkNewLogger(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewLogger()
	}
}

// BenchmarkRealWorldScenario 模拟真实场景的混合日志
func BenchmarkRealWorldScenario(b *testing.B) {
	logger := setupLogger()
	logger.SetFlag(FlagStackEnabled)
	logger.SetLevel(LevelInfo | LevelWarn | LevelError) // 过滤掉 Debug

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.I(benchTag, "Server started")
		logger.D(benchTag, "Debug information") // 会被过滤
		logger.W(benchTag, "Warning: high memory usage")
		logger.If(benchTag, "Request processed in %d ms", i%100)
		if i%10 == 0 {
			logger.E(benchTag, "Error occurred")
		}
	}
}
