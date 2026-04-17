package pretty_log

type (
	Flag  int // log flag
	Level int // log level
)

const (
	FlagClear        Flag = 0x00
	FlagColorEnabled Flag = 0x01 // enable color
	FlagStackEnabled Flag = 0x02 // enable stack info

	LevelInfo  Level = 0x01
	LevelDebug Level = 0x02
	LevelWarn  Level = 0x04
	LevelError Level = 0x08
	LevelFatal Level = 0x10
	LevelPanic Level = 0x20
)
