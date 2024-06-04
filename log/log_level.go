package log

type LogLevel int

const (
	LOG_LEVEL_DEBUG LogLevel = iota
	LOG_LEVEL_INFO
	LOG_LEVEL_WARNING
	LOG_LEVEL_ERROR
	LOG_LEVEL_CRITICAL
)
