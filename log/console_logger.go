package log

import (
	"time"
)

type consoleLogger struct{}

func (c consoleLogger) Handle(level LogLevel, msg string, details CallerDetails) error {
	logLevelToColor(level).Printf(
		"%s - [%s] %s#%d - %s\n",
		time.Now().Local().Format("2006-01-02 15:04:05"),
		logLevelToString(level),
		details.MethodName,
		details.LineNumber,
		msg,
	)
	return nil
}

// Register a log handler to print log messages to stdout.
func RegisterConsoleLogger() {
	RegisterHandler(consoleLogger{})
}

func logLevelToString(level LogLevel) string {
	switch level {
	case LOG_LEVEL_DEBUG:
		return "DEBUG"
	case LOG_LEVEL_INFO:
		return "INFO"
	case LOG_LEVEL_WARNING:
		return "WARNING"
	case LOG_LEVEL_ERROR:
		return "ERROR"
	case LOG_LEVEL_CRITICAL:
		return "CRITICAL"
	}
	return ""
}
