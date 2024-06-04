package log

// Holds informations about the caller of the log method.
type CallerDetails struct {
	FileName   string
	LineNumber int
	MethodName string
}
