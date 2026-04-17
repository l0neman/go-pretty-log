package main

import (
	"fmt"
	"time"

	plog "github.com/l0neman/go-pretty-log"
)

// This example shows how to optimize log performance based on benchmark results.

func main() {
	fmt.Println("=== Pretty-Log Performance Optimization Example ===")
	fmt.Println()

	// 1. High performance mode - for high concurrency scenarios
	fmt.Println("1. High Performance Mode (disable all flags)")
	highPerfLogger := plog.NewLogger()
	highPerfLogger.SetFlag(plog.FlagClear)
	highPerfLogger.SetLevel(plog.LevelWarn | plog.LevelError)

	start := time.Now()
	for i := 0; i < 10000; i++ {
		highPerfLogger.I("perf", "This info will be filtered")
		highPerfLogger.W("perf", "This is a warning")
	}
	fmt.Printf("   Duration: %v\n\n", time.Since(start))

	// 2. Development mode - for development debugging
	fmt.Println("2. Development Debug Mode (enable color and stack info)")
	devLogger := plog.NewLogger()
	devLogger.SetFlag(plog.FlagColorEnabled | plog.FlagStackEnabled)

	start = time.Now()
	for i := 0; i < 10000; i++ {
		devLogger.I("dev", "Debug information")
	}
	fmt.Printf("   Duration: %v\n\n", time.Since(start))

	// 3. Production mode - balance performance and features
	fmt.Println("3. Production Mode (enable stack info, disable color)")
	prodLogger := plog.NewLogger()
	prodLogger.SetFlag(plog.FlagStackEnabled)
	prodLogger.SetLevel(plog.LevelInfo | plog.LevelWarn | plog.LevelError)

	start = time.Now()
	for i := 0; i < 10000; i++ {
		prodLogger.D("prod", "This debug will be filtered")
		prodLogger.I("prod", "Important information")
	}
	fmt.Printf("   Duration: %v\n\n", time.Since(start))

	// 4. Show performance comparison of different configurations
	fmt.Println("\n=== Performance Comparison Test ===")

	configs := []struct {
		name  string
		flag  plog.Flag
		level plog.Level
	}{
		{"Fastest", plog.FlagClear, plog.LevelError},
		{"Color Only", plog.FlagColorEnabled, plog.LevelInfo | plog.LevelDebug | plog.LevelWarn | plog.LevelError},
		{"Stack Only", plog.FlagStackEnabled, plog.LevelInfo | plog.LevelDebug | plog.LevelWarn | plog.LevelError},
		{"Full Feature", plog.FlagColorEnabled | plog.FlagStackEnabled, plog.LevelInfo | plog.LevelDebug | plog.LevelWarn | plog.LevelError},
	}

	const iterations = 10000
	for _, cfg := range configs {
		logger := plog.NewLogger()
		logger.SetFlag(cfg.flag)
		logger.SetLevel(cfg.level)

		start := time.Now()
		for i := 0; i < iterations; i++ {
			logger.I("bench", "Test message")
		}
		duration := time.Since(start)

		fmt.Printf("%-12s: %v (avg %v/op)\n",
			cfg.name,
			duration,
			duration/iterations)
	}

	// 5. Show performance advantages of log filtering
	fmt.Println("\n=== Log Filtering Performance Advantage ===")

	filterLogger := plog.NewLogger()
	filterLogger.SetFlag(plog.FlagStackEnabled)

	// test filtering scenario
	filterLogger.SetLevel(plog.LevelError)
	start = time.Now()
	for i := 0; i < 100000; i++ {
		filterLogger.I("filter", "This will be filtered")
		filterLogger.D("filter", "This will be filtered")
		filterLogger.W("filter", "This will be filtered")
	}
	filteredDuration := time.Since(start)

	// test non-filtering scenario
	filterLogger.SetLevel(plog.LevelInfo | plog.LevelDebug | plog.LevelWarn)
	start = time.Now()
	for i := 0; i < 100000; i++ {
		filterLogger.I("filter", "This will be logged")
		filterLogger.D("filter", "This will be logged")
		filterLogger.W("filter", "This will be logged")
	}
	nonFilteredDuration := time.Since(start)

	fmt.Printf("Filter 100000 logs duration:    %v\n", filteredDuration)
	fmt.Printf("No filter 100000 logs duration: %v\n", nonFilteredDuration)
	fmt.Printf("Performance improvement:        %.2fx\n", float64(nonFilteredDuration)/float64(filteredDuration))

	// 6. Real application scenario examples
	fmt.Println()
	fmt.Println("=== Real Application Scenario Suggestions ===")
	fmt.Print(`
Based on benchmark results, here are configuration suggestions for different scenarios:

1. High frequency logging (QPS > 10000):
   - Disable all flags
   - Set higher log level (only record Warn and Error)
   - Expected performance: ~780 ns/op

2. Development debugging:
   - Enable color and stack info
   - Record all levels
   - Expected performance: ~1500 ns/op

3. Production environment:
   - Enable stack info only (for troubleshooting)
   - Filter Debug level
   - Expected performance: ~1300 ns/op

4. Performance testing/stress testing:
   - Disable logging or only record errors
   - Use log level filtering
   - Expected performance: ~70 ns/op (when filtering)

Key optimization points:
- Stack info retrieval is the biggest overhead (~450 ns)
- Color processing has moderate overhead (~200 ns)
- Log filtering is very efficient (only 5% overhead)
- Excellent concurrency performance (no lock contention)
`)

	// 7. Best practices
	fmt.Println()
	fmt.Println("=== Best Practices ===")
	fmt.Print(`
1. Use different configurations in different environments:
   - Development: full features (for easy debugging)
   - Testing: enable stack info (for easy problem location)
   - Production: adjust based on actual needs

2. Use appropriate log levels:
   - Debug: detailed debug information
   - Info: important business information
   - Warn: warning information (needs attention but does not affect running)
   - Error: error information (needs immediate handling)

3. Avoid logging in performance critical paths:
   - Use if check to determine if logging is needed
   - Use log level filtering (almost no overhead)

4. Formatting suggestions:
   - Simple logs use I/D/W/E directly
   - Use If/Df/Wf/Ef for formatted logs
   - Complex formatting has little performance impact (~5%)
`)
}
