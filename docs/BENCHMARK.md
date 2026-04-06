# 基准测试文档

## 运行基准测试

### 运行所有基准测试

```shell
go test -bench=. -benchmem
```

### 运行特定基准测试

```shell
# 只运行 Info 级别相关的测试
go test -bench=BenchmarkLoggerInfo -benchmem

# 运行所有并发测试
go test -bench=Parallel -benchmem

# 运行 flag 组合测试
go test -bench=BenchmarkFlagCombinations -benchmem
```

### 增加运行时间以获得更准确的结果

```shell
go test -bench=. -benchmem -benchtime=10s
```

### 生成性能分析文件

```shell
# CPU 分析
go test -bench=. -cpuprofile=cpu.prof

# 内存分析
go test -bench=. -memprofile=mem.prof

# 查看分析结果
go tool pprof cpu.prof
go tool pprof mem.prof
```

### 对比测试结果

```shell
# 运行基准测试并保存结果
go test -bench=. -benchmem > old.txt

# 修改代码后再次运行
go test -bench=. -benchmem > new.txt

# 使用 benchcmp 或 benchstat 对比（需要先安装）
# go install golang.org/x/perf/cmd/benchstat@latest
benchstat old.txt new.txt
```

## 基准测试说明

### 核心功能测试

- **BenchmarkLoggerInfo/Debug/Warn/Error**: 测试不同日志级别的性能
- **BenchmarkLoggerInfof**: 测试格式化日志的性能
- **BenchmarkGlobalLoggerInfo**: 测试全局 logger 的性能

### Flag 配置测试

- **BenchmarkLoggerWithColorEnabled/Disabled**: 测试颜色开关对性能的影响
- **BenchmarkLoggerWithStackEnabled/Disabled**: 测试栈信息开关对性能的影响
- **BenchmarkLoggerWithAllFlagsEnabled**: 测试所有 flag 启用时的性能
- **BenchmarkFlagCombinations**: 测试不同 flag 组合的性能对比

### 过滤测试

- **BenchmarkLoggerFiltered**: 测试日志被过滤时的性能（预期非常快）
- **BenchmarkLoggerNotFiltered**: 测试日志未被过滤时的性能

### 格式化对比测试

- **BenchmarkFormattingComparison**: 对比非格式化、简单格式化和复杂格式化的性能
- **BenchmarkLoggerMultipleArgs**: 测试多参数日志的性能
- **BenchmarkLoggerComplexFormatting**: 测试复杂格式化的性能

### 并发测试

- **BenchmarkLoggerParallel**: 测试基本并发日志写入性能
- **BenchmarkLoggerParallelWithAllFlags**: 测试启用所有 flag 时的并发性能

### 组件测试

- **BenchmarkPrinterImpl**: 测试 Printer 的性能
- **BenchmarkPrinterImplWithColor**: 测试带颜色的 Printer 性能
- **BenchmarkPrinterImplWithStack**: 测试带栈信息的 Printer 性能
- **BenchmarkGetStackInfo**: 测试获取栈信息的性能开销
- **BenchmarkNewLogger**: 测试创建 Logger 的性能

### 真实场景测试

- **BenchmarkRealWorldScenario**: 模拟真实应用中的混合日志场景

## 性能优化建议

根据基准测试结果，可以得出以下优化建议：

1. **禁用不需要的功能**: 如果不需要颜色或栈信息，禁用相关 flag 可以提升约 50% 的性能
2. **使用日志级别过滤**: 在生产环境中适当提高日志级别，被过滤的日志开销非常小
3. **避免过度格式化**: 简单的日志输出比复杂格式化更快
4. **并发性能良好**: 该日志库在高并发场景下表现稳定

## 基准测试结果示例

```
BenchmarkLoggerInfo-16                    	  730578	      1576 ns/op	     936 B/op	      21 allocs/op
BenchmarkLoggerWithColorDisabled-16       	 1505474	       791.0 ns/op	     416 B/op	      13 allocs/op
BenchmarkLoggerFiltered-16                	17101891	        72.42 ns/op	      64 B/op	       2 allocs/op
BenchmarkGetStackInfo-16                  	 2556325	       443.7 ns/op	     296 B/op	       5 allocs/op
```

说明：
- 启用栈信息会增加约 400-450 ns 的开销
- 启用颜色会增加约 200 ns 的开销
- 日志过滤的开销极小（约 72 ns）
- 基本日志操作在 1500 ns 左右

## 自定义基准测试

如果需要测试特定场景，可以参考 `log_bench_test.go` 中的示例编写自定义基准测试：

```go
func BenchmarkCustomScenario(b *testing.B) {
    logger := setupLogger()
    // 配置你的 logger
    logger.SetFlag(FlagStackEnabled)
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        // 你的测试代码
        logger.I("tag", "message")
    }
}
```

## 注意事项

1. 基准测试使用 `mockPrinter` 或 `discardPrinter` 来避免实际 I/O 操作影响结果
2. 测试结果可能因系统负载、CPU 型号等因素而异
3. 建议多次运行取平均值以获得更可靠的结果
4. 在实际生产环境中，I/O 操作（如写文件）会显著影响性能

