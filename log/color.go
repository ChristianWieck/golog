package log

import (
	"github.com/fatih/color"
)

var colorMappings = map[LogLevel]*color.Color{
	LOG_LEVEL_DEBUG:    color.New(color.Faint),
	LOG_LEVEL_INFO:     color.New(color.FgWhite),
	LOG_LEVEL_WARNING:  color.New(color.FgYellow),
	LOG_LEVEL_ERROR:    color.New(color.FgHiRed),
	LOG_LEVEL_CRITICAL: color.New(color.FgRed),
}

func logLevelToColor(l LogLevel) *color.Color {
	return colorMappings[l]
}
