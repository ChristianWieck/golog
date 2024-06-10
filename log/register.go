package log

var logHandlers = []struct {
	LogHandler
	LogLevel
}{}

type LogHandler interface {
	Handle(LogLevel, string, CallerDetails) error
}

func RegisterHandler(handler LogHandler) {
	RegisterHandlerVerbosity(handler, LogLevelDebug)
}

func RegisterHandlerVerbosity(handler LogHandler, verbosity LogLevel) {
	logHandlers = append(logHandlers, struct {
		LogHandler
		LogLevel
	}{handler, verbosity})
}
