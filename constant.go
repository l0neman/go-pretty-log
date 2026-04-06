package pretty_log

type (
	Flag  int // 日志标记
	Level int // 日志级别
)

const (
	FlagClear        Flag = 0x00
	FlagColorEnabled Flag = 0x01 // 启用颜色
	FlagStackEnabled Flag = 0x02 // 启用栈信息

	LevelInfo  Level = 0x01
	LevelDebug Level = 0x02
	LevelWarn  Level = 0x04
	LevelError Level = 0x08
	LevelFatal Level = 0x10
	LevelPanic Level = 0x20
)
