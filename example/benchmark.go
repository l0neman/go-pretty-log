package main

import (
	"fmt"
	"time"

	plog "github.com/l0neman/go-pretty-log"
)

// 这个示例展示了如何根据基准测试结果优化日志性能

func main() {
	fmt.Println("=== Pretty-Log 性能优化示例 ===")
	fmt.Println()

	// 1. 高性能模式 - 适用于高并发场景
	fmt.Println("1. 高性能模式（禁用所有 flag）")
	highPerfLogger := plog.NewLogger()
	highPerfLogger.SetFlag(plog.FlagClear)
	highPerfLogger.SetLevel(plog.LevelWarn | plog.LevelError)

	start := time.Now()
	for i := 0; i < 10000; i++ {
		highPerfLogger.I("perf", "This info will be filtered")
		highPerfLogger.W("perf", "This is a warning")
	}
	fmt.Printf("   耗时: %v\n\n", time.Since(start))

	// 2. 开发模式 - 适用于开发调试
	fmt.Println("2. 开发调试模式（启用颜色和栈信息）")
	devLogger := plog.NewLogger()
	devLogger.SetFlag(plog.FlagColorEnabled | plog.FlagStackEnabled)

	start = time.Now()
	for i := 0; i < 10000; i++ {
		devLogger.I("dev", "Debug information")
	}
	fmt.Printf("   耗时: %v\n\n", time.Since(start))

	// 3. 生产模式 - 平衡性能和功能
	fmt.Println("3. 生产环境模式（启用栈信息，禁用颜色）")
	prodLogger := plog.NewLogger()
	prodLogger.SetFlag(plog.FlagStackEnabled)
	prodLogger.SetLevel(plog.LevelInfo | plog.LevelWarn | plog.LevelError)

	start = time.Now()
	for i := 0; i < 10000; i++ {
		prodLogger.D("prod", "This debug will be filtered")
		prodLogger.I("prod", "Important information")
	}
	fmt.Printf("   耗时: %v\n\n", time.Since(start))

	// 4. 展示不同配置的性能对比
	fmt.Println("\n=== 性能对比测试 ===")

	configs := []struct {
		name  string
		flag  plog.Flag
		level plog.Level
	}{
		{"最快配置", plog.FlagClear, plog.LevelError},
		{"仅颜色", plog.FlagColorEnabled, plog.LevelInfo | plog.LevelDebug | plog.LevelWarn | plog.LevelError},
		{"仅栈信息", plog.FlagStackEnabled, plog.LevelInfo | plog.LevelDebug | plog.LevelWarn | plog.LevelError},
		{"完整功能", plog.FlagColorEnabled | plog.FlagStackEnabled, plog.LevelInfo | plog.LevelDebug | plog.LevelWarn | plog.LevelError},
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

		fmt.Printf("%-12s: %v (平均 %v/op)\n",
			cfg.name,
			duration,
			duration/iterations)
	}

	// 5. 展示日志过滤的性能优势
	fmt.Println("\n=== 日志过滤性能优势 ===")

	filterLogger := plog.NewLogger()
	filterLogger.SetFlag(plog.FlagStackEnabled)

	// 测试过滤场景
	filterLogger.SetLevel(plog.LevelError) // 只记录错误
	start = time.Now()
	for i := 0; i < 100000; i++ {
		filterLogger.I("filter", "This will be filtered")
		filterLogger.D("filter", "This will be filtered")
		filterLogger.W("filter", "This will be filtered")
	}
	filteredDuration := time.Since(start)

	// 测试非过滤场景
	filterLogger.SetLevel(plog.LevelInfo | plog.LevelDebug | plog.LevelWarn)
	start = time.Now()
	for i := 0; i < 100000; i++ {
		filterLogger.I("filter", "This will be logged")
		filterLogger.D("filter", "This will be logged")
		filterLogger.W("filter", "This will be logged")
	}
	nonFilteredDuration := time.Since(start)

	fmt.Printf("过滤 100000 条日志耗时:    %v\n", filteredDuration)
	fmt.Printf("不过滤 100000 条日志耗时:  %v\n", nonFilteredDuration)
	fmt.Printf("性能提升:                  %.2fx\n", float64(nonFilteredDuration)/float64(filteredDuration))

	// 6. 实际应用场景示例
	fmt.Println()
	fmt.Println("=== 实际应用场景建议 ===")
	fmt.Print(`
根据基准测试结果，以下是不同场景的配置建议：

1. 高频日志场景（QPS > 10000）：
   - 禁用所有 flag
   - 设置较高的日志级别（只记录 Warn 和 Error）
   - 预期性能: ~780 ns/op

2. 开发调试场景：
   - 启用颜色和栈信息
   - 记录所有级别
   - 预期性能: ~1500 ns/op

3. 生产环境：
   - 只启用栈信息（方便排查问题）
   - 过滤 Debug 级别
   - 预期性能: ~1300 ns/op

4. 性能测试/压测场景：
   - 禁用日志或只记录错误
   - 使用日志级别过滤
   - 预期性能: ~70 ns/op（过滤时）

关键优化点：
- 栈信息获取是最大开销（~450 ns）
- 颜色处理有适度开销（~200 ns）
- 日志过滤非常高效（仅 5% 开销）
- 并发性能优秀（无锁竞争）
`)

	// 7. 最佳实践
	fmt.Println()
	fmt.Println("=== 最佳实践 ===")
	fmt.Print(`
1. 在不同环境使用不同配置：
   - 开发环境: 全功能（方便调试）
   - 测试环境: 启用栈信息（方便问题定位）
   - 生产环境: 根据实际需求调整

2. 使用合适的日志级别：
   - Debug: 详细的调试信息
   - Info:  重要的业务信息
   - Warn:  警告信息（需要关注但不影响运行）
   - Error: 错误信息（需要立即处理）

3. 避免在性能关键路径打日志：
   - 可以使用 if 判断是否需要记录
   - 利用日志级别过滤（几乎无开销）

4. 格式化建议：
   - 简单日志直接使用 I/D/W/E
   - 需要格式化时使用 If/Df/Wf/Ef
   - 复杂格式化的性能影响很小（~5%）
`)
}
