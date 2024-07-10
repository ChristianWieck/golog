package log

import (
	"time"
)

type consoleLogger struct{}

func (c consoleLogger) Handle(level LogLevel, msg string, details LogDetails) error {
	logLevelToColor(level).Printf(
		"%s - [%s] %s - %s\n",
		time.Now().Local().Format("2006-01-02 15:04:05"),
		logLevelToString(level),
		details.LoggerName,
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
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarning:
		return "WARNING"
	case LogLevelError:
		return "ERROR"
	case LogLevelCritical:
		return "CRITICAL"
	}
	return ""
}
