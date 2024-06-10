package log

import (
	"github.com/fatih/color"
)

var colorMappings = map[LogLevel]*color.Color{
	LogLevelDebug:    color.New(color.Faint),
	LogLevelInfo:     color.New(color.FgWhite),
	LogLevelWarning:  color.New(color.FgYellow),
	LogLevelError:    color.New(color.FgHiRed),
	LogLevelCritical: color.New(color.FgRed),
}

func logLevelToColor(l LogLevel) *color.Color {
	return colorMappings[l]
}
