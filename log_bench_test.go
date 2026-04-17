package pretty_log

import (
	"io"
	"os"
	"testing"
	"time"
)

// mockPrinter A mock outputter for benchmarking that discards all output
type mockPrinter struct{}

func (m *mockPrinter) Print(time time.Time, level Level, logTag, logText string, pid int, colorful bool, stackInfo string) {
	// Do not make any output, simulate the fastest printer
}

// discardPrinter is a printer that uses io.Discard
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

// setupLogger creates a logger for benchmark testing
func setupLogger() Logger {
	logger := NewLogger()
	logger.SetPrinter(newDiscardPrinter())
	return logger
}

// BenchmarkLoggerInfo tests Info level log performance
func BenchmarkLoggerInfo(b *testing.B) {
	logger := setupLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.I(benchTag, benchMessage)
	}
}

// BenchmarkLoggerInfof tests formatted Info log performance
func BenchmarkLoggerInfof(b *testing.B) {
	logger := setupLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.If(benchTag, "This is a %s test message %d", "benchmark", i)
	}
}

// BenchmarkLoggerDebug tests Debug level log performance
func BenchmarkLoggerDebug(b *testing.B) {
	logger := setupLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.D(benchTag, benchMessage)
	}
}

// BenchmarkLoggerWarn tests Warn level log performance
func BenchmarkLoggerWarn(b *testing.B) {
	logger := setupLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.W(benchTag, benchMessage)
	}
}

// BenchmarkLoggerError tests Error level log performance
func BenchmarkLoggerError(b *testing.B) {
	logger := setupLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.E(benchTag, benchMessage)
	}
}

// BenchmarkGlobalLoggerInfo tests global logger performance
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

// BenchmarkLoggerWithColorEnabled tests performance with color enabled
func BenchmarkLoggerWithColorEnabled(b *testing.B) {
	logger := setupLogger()
	logger.SetFlag(FlagColorEnabled)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.I(benchTag, benchMessage)
	}
}

// BenchmarkLoggerWithColorDisabled tests performance with color disabled
func BenchmarkLoggerWithColorDisabled(b *testing.B) {
	logger := setupLogger()
	logger.SetFlag(FlagClear)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.I(benchTag, benchMessage)
	}
}

// BenchmarkLoggerWithStackEnabled tests performance with stack info enabled
func BenchmarkLoggerWithStackEnabled(b *testing.B) {
	logger := setupLogger()
	logger.SetFlag(FlagStackEnabled)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.I(benchTag, benchMessage)
	}
}

// BenchmarkLoggerWithStackDisabled tests performance with stack info disabled
func BenchmarkLoggerWithStackDisabled(b *testing.B) {
	logger := setupLogger()
	logger.SetFlag(FlagClear)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.I(benchTag, benchMessage)
	}
}

// BenchmarkLoggerWithAllFlagsEnabled tests performance with all flags enabled
func BenchmarkLoggerWithAllFlagsEnabled(b *testing.B) {
	logger := setupLogger()
	logger.SetFlag(FlagColorEnabled | FlagStackEnabled)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.I(benchTag, benchMessage)
	}
}

// BenchmarkLoggerFiltered tests performance when logs are filtered
func BenchmarkLoggerFiltered(b *testing.B) {
	logger := setupLogger()
	logger.SetLevel(LevelError)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.I(benchTag, benchMessage)
	}
}

// BenchmarkLoggerNotFiltered tests performance when logs are not filtered
func BenchmarkLoggerNotFiltered(b *testing.B) {
	logger := setupLogger()
	logger.SetLevel(LevelInfo)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.I(benchTag, benchMessage)
	}
}

// BenchmarkLoggerMultipleArgs tests performance of multi-argument logs
func BenchmarkLoggerMultipleArgs(b *testing.B) {
	logger := setupLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.I(benchTag, "arg1", "arg2", "arg3", 123, 456.789)
	}
}

// BenchmarkLoggerComplexFormatting tests performance of complex formatting
func BenchmarkLoggerComplexFormatting(b *testing.B) {
	logger := setupLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.If(benchTag, "User %s logged in at %d with score %.2f from %s", "john_doe", i, 95.5, "192.168.1.1")
	}
}

// BenchmarkLoggerParallel tests concurrent log write performance
func BenchmarkLoggerParallel(b *testing.B) {
	logger := setupLogger()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.I(benchTag, benchMessage)
		}
	})
}

// BenchmarkLoggerParallelWithAllFlags tests concurrent performance with all flags enabled
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

// BenchmarkDifferentLevels compares performance of different log levels
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

// BenchmarkFormattingComparison compares formatted and non-formatted performance
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

// BenchmarkFlagCombinations tests performance of different flag combinations
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

// BenchmarkPrinterImpl tests default Printer performance
func BenchmarkPrinterImpl(b *testing.B) {
	printer := newDiscardPrinter()
	now := time.Now()
	pid := os.Getpid()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		printer.Print(now, LevelInfo, benchTag, benchMessage, pid, false, "")
	}
}

// BenchmarkPrinterImplWithColor tests Printer performance with color
func BenchmarkPrinterImplWithColor(b *testing.B) {
	printer := newDiscardPrinter()
	now := time.Now()
	pid := os.Getpid()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		printer.Print(now, LevelInfo, benchTag, benchMessage, pid, true, "")
	}
}

// BenchmarkPrinterImplWithStack tests Printer performance with stack info
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

// BenchmarkGetStackInfo tests performance of getting stack info
func BenchmarkGetStackInfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = getStackInfo(baseStackOffset)
	}
}

// BenchmarkNewLogger tests performance of creating new Logger
func BenchmarkNewLogger(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewLogger()
	}
}

// BenchmarkRealWorldScenario simulates mixed logs in real scenarios
func BenchmarkRealWorldScenario(b *testing.B) {
	logger := setupLogger()
	logger.SetFlag(FlagStackEnabled)
	logger.SetLevel(LevelInfo | LevelWarn | LevelError)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.I(benchTag, "Server started")
		logger.D(benchTag, "Debug information")
		logger.W(benchTag, "Warning: high memory usage")
		logger.If(benchTag, "Request processed in %d ms", i%100)
		if i%10 == 0 {
			logger.E(benchTag, "Error occurred")
		}
	}
}
