package log

import (
	"os"
	"runtime"
)

func logInternal(level LogLevel, msg string) {
	// Recover information about the caller
	pc, file, no, ok := runtime.Caller(2) // We need to go two steps back
	details := runtime.FuncForPC(pc)
	var logDetails CallerDetails
	if ok {
		logDetails = CallerDetails{FileName: file, LineNumber: no, MethodName: details.Name()}
	}

	for _, v := range logHandlers {
		if level >= v.LogLevel {
			v.Handle(level, msg, logDetails)
		}
	}
}

func Log(level LogLevel, msg string) {
	logInternal(level, msg)
}

// Log a message with "Debug" level
func Debug(msg string) { logInternal(LOG_LEVEL_DEBUG, msg) }

// Log a message with "Info" level
func Info(msg string) { logInternal(LOG_LEVEL_INFO, msg) }

// Log a message with "Warning" level
func Warning(msg string) { logInternal(LOG_LEVEL_WARNING, msg) }

// Log a message with "Error" level
func Error(msg string) { logInternal(LOG_LEVEL_ERROR, msg) }

// Log a message with "Critical" level. This will also interrupt the program once
// the message is logged with an exit code of 1.
func Critical(msg string) {
	logInternal(LOG_LEVEL_CRITICAL, msg)
	os.Exit(1)
}
