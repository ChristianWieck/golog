package log

import (
	"fmt"
	"runtime"
)

type Logger struct {
	name string
}

func Create(name string) *Logger {
	return &Logger{name: name}
}

func (l *Logger) logInternal(level LogLevel, format string, a ...any) {
	var logDetails LogDetails
	msg := ""
	msgFormatted := false

	handlerLock.Lock()
	defer handlerLock.Unlock()
	for _, v := range logHandlers {
		if level >= v.LogLevel {
			if !msgFormatted {
				msg = fmt.Sprintf(format, a...)

				// Also collect the caller information
				pc, file, no, ok := runtime.Caller(2) // We need to go two steps back
				details := runtime.FuncForPC(pc)

				if ok {
					logDetails = LogDetails{
						LoggerName: l.name,
						FileName:   file,
						LineNumber: no,
						MethodName: details.Name(),
					}
				}
				msgFormatted = true
			}
			v.LogHandler.Handle(level, msg, logDetails)
		}
	}
}

func (l *Logger) Log(level LogLevel, format string, a ...any) { l.logInternal(level, format, a...) }

// Log a message with "Debug" level
func (l *Logger) Debug(format string, a ...any) { l.logInternal(LogLevelDebug, format, a...) }

// Log a message with "Info" level
func (l *Logger) Info(format string, a ...any) { l.logInternal(LogLevelInfo, format, a...) }

// Log a message with "Warning" level
func (l *Logger) Warning(format string, a ...any) { l.logInternal(LogLevelWarning, format, a...) }

// Log a message with "Error" level
func (l *Logger) Error(format string, a ...any) { l.logInternal(LogLevelError, format, a...) }

// Log a message with "Critical" level. This will also interrupt the program once
// the message is logged with an exit code of 1.
func (l *Logger) Critical(format string, a ...any) {
	l.logInternal(LogLevelCritical, format, a...)
	panic(fmt.Sprintf(format, a...))
}
