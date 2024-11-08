package log

import (
	"testing"
	"time"
)

type testLogger struct {
	t *testing.T
}

func (t testLogger) Handle(level LogLevel, msg string, details LogDetails) error {
	t.t.Logf(
		"%s - [%s] %s - %s\n",
		time.Now().Local().Format("2006-01-02 15:04:05.000"),
		logLevelToString(level),
		details.LoggerName,
		msg,
	)
	return nil
}

// Register a log handler to print log messages to test output.
func RegisterTestLogger(t *testing.T) {
	RegisterHandler(
		testLogger{
			t: t,
		},
	)
}

// Register a log handler to print log messages with the given level or higher to test output.
func RegisterTestLoggerVerbosity(t *testing.T, level LogLevel) {
	RegisterHandlerVerbosity(
		testLogger{
			t: t,
		},
		level,
	)
}
