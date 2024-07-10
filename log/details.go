package log

// Holds informations about the caller of the log method.
type LogDetails struct {
	LoggerName string
	FileName   string
	LineNumber int
	MethodName string
}
