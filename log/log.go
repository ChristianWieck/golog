package log

import (
	"fmt"
	"runtime"
)

func logInternal(level LogLevel, format string, a ...any) {
	// Recover information about the caller
	pc, file, no, ok := runtime.Caller(2) // We need to go two steps back
	details := runtime.FuncForPC(pc)
	var logDetails CallerDetails
	if ok {
		logDetails = CallerDetails{FileName: file, LineNumber: no, MethodName: details.Name()}
	}

	msg := ""
	msgFormatted := false
	handlerLock.Lock()
	defer handlerLock.Unlock()
	for _, v := range logHandlers {
		if level >= v.LogLevel {
			if !msgFormatted {
				msg = fmt.Sprintf(format, a...)
				msgFormatted = true
			}
			v.Handle(level, msg, logDetails)
		}
	}
}

func Log(level LogLevel, format string, a ...any) { logInternal(level, format, a...) }

// Log a message with "Debug" level
func Debug(format string, a ...any) { logInternal(LogLevelDebug, format, a...) }

// Log a message with "Info" level
func Info(format string, a ...any) { logInternal(LogLevelInfo, format, a...) }

// Log a message with "Warning" level
func Warning(format string, a ...any) { logInternal(LogLevelWarning, format, a...) }

// Log a message with "Error" level
func Error(format string, a ...any) { logInternal(LogLevelError, format, a...) }

// Log a message with "Critical" level. This will also interrupt the program once
// the message is logged with an exit code of 1.
func Critical(format string, a ...any) {
	logInternal(LogLevelCritical, format, a...)
	panic(fmt.Sprintf(format, a...))
}
